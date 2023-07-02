package main

import (
	"log"

	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/config"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
