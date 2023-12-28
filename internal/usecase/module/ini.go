package uc_module

import (
	"git.legchelife.ru/root/template/internal/repo/db"
	"git.legchelife.ru/root/template/internal/usecase/module/user"
)

type Repo struct {
	repoDB *repo_db.RepoDB
	User   *uc_user.UseCase
}

func New(repoDB *repo_db.RepoDB) *Repo {

	user := uc_user.New(repoDB)

	return &Repo{
		repoDB: repoDB,
		User:   user,
	}
}
