package uc

import (
	"git.legchelife.ru/root/template/internal/repo/service"
	uc_module "git.legchelife.ru/root/template/internal/usecase/module"
)

type Repo struct {
	repo   *service.Repo
	Module *uc_module.Repo
}

// New -.
func New(repo *service.Repo) *Repo {
	module := uc_module.New(repo)

	return &Repo{
		repo:   repo,
		Module: module,
	}
}
