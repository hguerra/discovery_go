package main

import (
	"config/pkg/config"
	"flag"
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

	// log.Println("---------------------------------------")
	// defaultAppEnv := flag.String("env", "development", "App environment")
	// flag.Parse()
	// log.Println(config.Getenv("APP_ENV2", *defaultAppEnv))

	log.Println("---------------------------------------")
	var appEnv string
	const appEnvKey = "env"
	const defaultAppEnvValue = "development"
	const defaultAppEnvDescription = "App environment"

	webCmd := flag.NewFlagSet("web", flag.ExitOnError)
	workerCmd := flag.NewFlagSet("worker", flag.ExitOnError)
	dbCmd := flag.NewFlagSet("db", flag.ExitOnError)
	if len(os.Args) < 2 {
		fmt.Println("subcommand is mandatory")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "web":
		webCmd.StringVar(&appEnv, appEnvKey, defaultAppEnvValue, defaultAppEnvDescription)
		host := webCmd.String("host", "localhost", "HTTP server host")
		webCmd.Parse(os.Args[2:])

		fmt.Println("	host:", *host)
	case "worker":
		workerCmd.StringVar(&appEnv, appEnvKey, defaultAppEnvValue, defaultAppEnvDescription)
		count := workerCmd.Int("count", 1, "Number of workers")
		workerCmd.Parse(os.Args[2:])

		fmt.Println("	count:", *count)
	case "db":
		dbCmd.StringVar(&appEnv, appEnvKey, defaultAppEnvValue, defaultAppEnvDescription)
		dbCmd.Parse(os.Args[2:])
	default:
		fmt.Println("unexpected subcommand")
		os.Exit(1)
	}

	cfg2, err := config.Load2(fmt.Sprintf("./configs/config.%s.json", config.Getenv("APP_ENV3", appEnv)))
	if err != nil {
		panic(err)
	}

	log.Println(config.Getenv("APP_ENV3", appEnv))
	log.Println(cfg2.Name)
	log.Println(cfg2.Version)
	log.Println(cfg2.A.B)
	log.Println(cfg2.Endpoints[0].Endpoint)
}
