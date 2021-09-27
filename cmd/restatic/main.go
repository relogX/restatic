package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/relogHQ/restatic/config"
)

func main() {
	flag.Parse()
	log.Infof("serving directory: %s", config.Directory)
	log.Infof("starting http server at port %d", config.Port)
}
