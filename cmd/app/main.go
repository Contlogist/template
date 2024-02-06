package main

import (
	"git.legchelife.ru/root/template/config"
	"git.legchelife.ru/root/template/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	app.Run(cfg)
}
