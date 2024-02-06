// Package app configures and runs application.
package app

import (
	"fmt"
	"git.legchelife.ru/root/template/internal/repo"
	"git.legchelife.ru/root/template/internal/usecase"
	natsconnect "git.legchelife.ru/root/template/pkg/nats"
	"git.legchelife.ru/root/template/pkg/upper"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"git.legchelife.ru/root/template/config"
	"git.legchelife.ru/root/template/internal/controller/http/v1"
	"git.legchelife.ru/root/template/pkg/httpserver"
	"git.legchelife.ru/root/template/pkg/logger"
)

const (
	dev  = "dev"
	prod = "prod"
	test = "test"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New()

	// DB
	dbClient, err := upper.NewPostgres(cfg.PG.URL)
	if err != nil {
		logrus.Error(err)
	}

	// Repository
	repo := repo.New(&dbClient)

	// Use case
	useCase := uc.New(repo)

	//NATS Server
	server, err := natsconnect.Connect(cfg.Nats.URL)
	if err != nil {
		l.Error(fmt.Errorf("app - Run - natsConnect.Connect - params: %w", err), logger.Data{Module: "app", Method: "Run", Action: logger.Parse, Params: map[string]interface{}{"signal": err}})
		return
	}
	natsconnect.New(server, *useCase)

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
		l.Info("app - Run - signal - params:", logger.Data{Module: "app", Method: "Run", Action: logger.Parse, Params: map[string]interface{}{"signal": s}})
	case err = <-httpServer.Notify():
		l.Error(err, logger.Data{Module: "app", Method: "Run", Action: logger.Parse, Params: map[string]interface{}{"signal": err}})
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown - params: %w", err), logger.Data{Module: "app", Method: "Run", Action: logger.Parse, Params: map[string]interface{}{"signal": err}})
	}
}
