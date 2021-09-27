package server

import (
	"fmt"
	"net/http"

	"github.com/relogHQ/restatic/config"
	log "github.com/sirupsen/logrus"
)

func Run() {
	log.Infof("server listening on http://localhost:%d", config.Port)

	fs := http.FileServer(http.Dir(config.Directory))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), fs); err != nil {
		log.Fatal(err)
	}
}
