package uc

import (
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/repo/service"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/usecase/module"
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
