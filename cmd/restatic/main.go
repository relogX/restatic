package main

import (
	"flag"

	"github.com/relogHQ/restatic/server"
)

func main() {
	flag.Parse()
	server.Run()
}
