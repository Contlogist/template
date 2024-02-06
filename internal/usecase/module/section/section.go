package uc_section

import (
	"git.legchelife.ru/root/template/internal/repo"
	rp_section "git.legchelife.ru/root/template/internal/repo/section"
	"git.legchelife.ru/root/template/pkg/models/context"
)

type Section struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Section {
	return &Section{
		repo: repo,
	}
}

// GetList получает список моделей Section из репозитория.
func (uc *Section) GetList(ctx *context.Base) ([]rp_section.Section, error) {
	//ctx.SetTimeout(3)

	re, err := uc.repo.SectionRepo.GetList(ctx)
	if err != nil {
		return nil, err
	} else {
		return re, nil
	}
}
