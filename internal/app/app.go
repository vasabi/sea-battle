package app

import (
	log "github.com/sirupsen/logrus"
	"sea-battle/api/http_server"
	"sea-battle/internal/pkg/utils"
)

// Run - initializing the configuration and launching the application.
func Run() {
	if err := utils.InitEnvironmentSettings(); err != nil {
		log.Fatalf("cannot init settings", err)
	}

	http_server.RunServer()
}
