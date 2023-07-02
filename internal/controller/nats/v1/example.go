package natsrouter

import (
	"errors"
	uc "git.legchelife.ru/root/template/internal/usecase"
	"git.legchelife.ru/root/template/pkg/models/rechan"
)

type exampleRoutes struct {
	usecase uc.Repo
}

func newExampleRoutes(reChan chan rechan.Nats, dp DecodePayload, usecase uc.Repo) {
	r := &exampleRoutes{usecase}
	switch dp.Method {
	case "GET":
		r.getExample(reChan)
	case "POST":
		r.postExample(reChan)
	default:
		err := errors.New("method not found")
		reChan <- rechan.Nats{
			Error: err,
		}
	}
}

func (r *exampleRoutes) getExample(reChan chan rechan.Nats) {
	reChan <- rechan.Nats{
		Completed: true,
	}
}

func (r *exampleRoutes) postExample(reChan chan rechan.Nats) {
	reChan <- rechan.Nats{
		Completed: true,
	}
}
