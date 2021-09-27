package config

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var Port int
var Directory string
var LogLevel string

func initFlags() {
	flag.IntVar(&Port, "p", 5030, "port")
	flag.StringVar(&Directory, "d", ".", "directory")
	flag.StringVar(&LogLevel, "log-level", "info", "log level")
}

func initLog() {
	switch LogLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func init() {
	initFlags()
	initLog()
}
