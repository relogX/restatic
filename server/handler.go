package server

import "net/http"

type fsHandler struct{}

func (f fsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("ok"))
}

func NewFSHandler() *fsHandler {
	return &fsHandler{}
}
