package uc_module_component

import (
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/repo/service"
)

type Repo struct {
	repo *service.Repo
}

func New(repo *service.Repo) *Repo {
	return &Repo{
		repo: repo,
	}
}
