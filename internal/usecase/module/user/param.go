package uc_user

import (
	"git.legchelife.ru/root/template/internal/repo"
	"git.legchelife.ru/root/template/internal/repo/user"
	"git.legchelife.ru/root/template/pkg/models/context"
)

type Params struct {
	repo *repo.Repo
}

func NewParams(repo *repo.Repo) *Params {
	return &Params{
		repo: repo,
	}
}

// GetList - метод получения списка параметров пользователя.
func (uc *Params) GetList(ctx *context.Base, id int) ([]user.Param, error) {
	//ctx.SetTimeout(3)

	//re, err := uc.repo.User.Params(uc.repo.DB).GetList(ctx, id)
	re, err := uc.repo.UserRepo.Param.GetList(ctx, id)
	if err != nil {
		return nil, err
	} else {
		return re, nil
	}
}
