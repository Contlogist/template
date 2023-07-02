package uc_module

import (
	"git.legchelife.ru/root/template/internal/repo/service"
	uc_module_component "git.legchelife.ru/root/template/internal/usecase/module/component"
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
