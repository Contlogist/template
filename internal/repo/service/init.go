package service

import (
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/ent"
	service_module "git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/repo/service/module"
)

type Repo struct {
	Module *service_module.Repo
}

func New(db *ent.Client) *Repo {

	mod := service_module.New(db)

	return &Repo{
		mod,
	}
}
