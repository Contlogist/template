package uc_user

import (
	"git.legchelife.ru/root/template/internal/repo/db"
)

type UseCase struct {
	repo *repo_db.RepoDB
}

func New(repo *repo_db.RepoDB) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
