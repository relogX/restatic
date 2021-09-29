package server

import (
	"fmt"
	"net/http"

	"github.com/relogHQ/restatic/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	fmt.Println(`
	██████  ███████ ███████ ████████  █████  ████████ ██  ██████ 
	██   ██ ██      ██         ██    ██   ██    ██    ██ ██      
	██████  █████   ███████    ██    ███████    ██    ██ ██      
	██   ██ ██           ██    ██    ██   ██    ██    ██ ██      
	██   ██ ███████ ███████    ██    ██   ██    ██    ██  ██████ 
	
	by Relog - https://relog.in
	`)
}

func Run() {
	addr := fmt.Sprintf(":%d", config.Port)
	handler := NewFSHandler()

	log.Infof("server listening on %s", addr)
	log.Infof("=========================")

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
