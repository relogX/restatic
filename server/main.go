package server

import (
	"fmt"
	"net/http"

	"github.com/relogHQ/restatic/config"
	log "github.com/sirupsen/logrus"
)

func Run() {
	log.Infof("server listening on http://localhost:%d", config.Port)

	addr := fmt.Sprintf(":%d", config.Port)
	handler := NewFSHandler()
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
