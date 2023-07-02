package uc_module

import (
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/repo/service"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/usecase/module/component"
)

type Repo struct {
	repo      *service.Repo
	Component *uc_module_component.Repo
}

func New(repo *service.Repo) *Repo {

	component := uc_module_component.New(repo)

	return &Repo{
		repo:      repo,
		Component: component,
	}
}
