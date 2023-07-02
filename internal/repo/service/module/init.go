package service_module

import (
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/ent"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/repo/service/module/component"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/pkg/models/context"
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
