package web

import (
	"log"

	"github.com/hguerra/discovery_go/modules/config/configviper/pkg/config"
)

func NewServer(cfg *config.Config) {
	address := cfg.GetString("SERVER_ADDRESS")
	log.Printf("Listening and serving HTTP on %s", address)

	databaseURL := cfg.GetString("DATABASE_URL")
	log.Printf("PostgreSQL data source '%s'", databaseURL)
}
