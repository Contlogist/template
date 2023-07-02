package service_module

import (
	"git.legchelife.ru/root/template/ent"
	"git.legchelife.ru/root/template/internal/repo/service/module/component"
	"git.legchelife.ru/root/template/pkg/models/context"
)

type Repo struct {
	db        *ent.Client
	Component *service_module_component.Component
}

func New(db *ent.Client) *Repo {

	component := service_module_component.New(db)

	return &Repo{
		db,
		component,
	}
}

type (
	Category interface {
		GetTest(ctx context.Base) (bool, error)
	}
)
