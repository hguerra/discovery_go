package main

import (
	"config/pkg/config"
	"fmt"
	"log"
	"os"
)

func main() {
	os.Setenv("APP_ENV", "production")

	cfg, err := config.Load(fmt.Sprintf("./configs/config.%s.json", config.GetEnv()))
	if err != nil {
		panic(err)
	}

	log.Println(cfg.GetString("name"))
	log.Println(cfg.GetFloat64("version"))
	log.Println(cfg.GetString("a.b"))
	log.Println(cfg.GetString("endpoints.0.endpoint"))
}
