package uc

import (
	"git.legchelife.ru/root/template/internal/repo/db"
	"git.legchelife.ru/root/template/internal/usecase/module"
)

type Repo struct {
	repo   *repo_db.RepoDB
	Module *uc_module.Repo
}

// New -.
func New(repo *repo_db.RepoDB) *Repo {
	module := uc_module.New(repo)

	return &Repo{
		repo:   repo,
		Module: module,
	}
}
