package service

import (
	"git.legchelife.ru/root/template/ent"
	service_module "git.legchelife.ru/root/template/internal/repo/service/module"
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
