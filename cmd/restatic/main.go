package main

import (
	"flag"
	"path/filepath"

	"github.com/relogHQ/restatic/config"
	"github.com/relogHQ/restatic/server"
)

func main() {
	flag.Parse()
	if bPath, err := filepath.Abs(config.Directory); err == nil {
		config.Directory = bPath
	}
	server.Run()
}
