package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/relogHQ/restatic/config"
)

type fsHandler struct{}
type dirlist struct {
	Files   []*fInfo
	DirInfo *dInfo
}
type fInfo struct {
	Name    string
	Mode    string
	ModTime string
	Size    string
	Path    string
	IsDir   bool
}
type dInfo struct {
	Path string
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func toFInfo(entry os.DirEntry, pwd string) *fInfo {
	info, err := entry.Info()
	if err != nil {
		return nil
	}

	path, err := filepath.Rel(config.Directory, pwd)
	if err != nil {
		return nil
	}

	return &fInfo{
		Name:    entry.Name(),
		Mode:    info.Mode().String(),
		ModTime: info.ModTime().Format(time.RFC1123),
		Size:    ByteCountIEC(info.Size()),
		Path:    path,
		IsDir:   entry.Type().IsDir(),
	}
}

func toFInfos(infos []os.DirEntry, pwd string) []*fInfo {
	fInfos := make([]*fInfo, len(infos))
	for i, info := range infos {
		fInfos[i] = toFInfo(info, pwd)
	}
	return fInfos
}

func toDInfo(info os.FileInfo, pwd string) *dInfo {
	path, err := filepath.Rel(config.Directory, pwd)
	if err != nil {
		return nil
	}

	return &dInfo{
		Path: path,
	}
}

func write500(w http.ResponseWriter) {
	http.Error(w, http.StatusText(500), 500)
}

func writeDirectory(w http.ResponseWriter, path string, dirInfo os.FileInfo) {
	tmpl := template.Must(template.ParseFiles("templates/dir.html", "templates/layout.html"))

	files, err := os.ReadDir(path)
	if err != nil {
		write500(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, dirlist{
		Files:   toFInfos(files, path),
		DirInfo: toDInfo(dirInfo, path),
	})
}

func writeFile(w http.ResponseWriter, info os.FileInfo) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("file"))
}

func (f fsHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	baseDir := config.Directory
	if _, err := os.Stat(baseDir); err != nil {
		log.Fatal(err)
	}

	path := path.Clean(path.Join(baseDir, request.URL.Path))

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		http.NotFound(w, request)
		return
	}

	if err != nil {
		write500(w)
		return
	}

	if info.IsDir() {
		writeDirectory(w, path, info)
		return
	}

	writeFile(w, info)
}

func NewFSHandler() *fsHandler {
	return &fsHandler{}
}
