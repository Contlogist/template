package uc_module_component

import (
	"git.legchelife.ru/root/template/internal/repo/service"
)

type Repo struct {
	repo *service.Repo
}

func New(repo *service.Repo) *Repo {
	return &Repo{
		repo: repo,
	}
}
