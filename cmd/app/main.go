package main

import (
	"log"

	"git.legchelife.ru/root/template/config"
	"git.legchelife.ru/root/template/internal/app"
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
