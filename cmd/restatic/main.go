package main

import (
	"flag"
	"path/filepath"

	"github.com/relogX/restatic/config"
	"github.com/relogX/restatic/server"
)

func main() {
	flag.Parse()
	if bPath, err := filepath.Abs(config.Directory); err == nil {
		config.Directory = bPath
	}
	server.Run()
}
