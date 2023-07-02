// Package app configures and runs application.
package app

import (
	"fmt"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/repo/service"
	uc "git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/usecase"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/pkg/ent"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/config"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/controller/http/v1"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/pkg/httpserver"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/pkg/logger"
)

const (
	dev  = "dev"
	prod = "prod"
	test = "test"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	migration := true
	switch cfg.App.Build {
	case prod:
		migration = false
	case dev:
		migration = true
	}
	entClient, err := ent.NewPostgresClient(cfg.EntPG, migration)

	// Repository
	repo := service.New(entClient)

	// Use case
	useCase := uc.New(repo)

	//NATS Server
	//conn, err := natsconnect.Connect(cfg.Nats.URL)
	//natsconnect.New(conn, useCase)

	// HTTP Server
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	v1.NewRouter(handler, l, *useCase)
	httpServer := httpserver.New(handler, cfg.App.Build, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
