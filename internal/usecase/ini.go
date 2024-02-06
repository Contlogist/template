package uc

import (
	"git.legchelife.ru/root/template/internal/repo"
	"git.legchelife.ru/root/template/internal/usecase/module/section"
	"git.legchelife.ru/root/template/internal/usecase/module/user"
)

type Repo struct {
	repo        *repo.Repo
	UserRepo    *uc_user.UserRepo
	SectionRepo *uc_section.Section
}

// New -.
func New(repo *repo.Repo) *Repo {

	return &Repo{
		repo:        repo,
		UserRepo:    uc_user.New(repo),
		SectionRepo: uc_section.New(repo),
	}
}
