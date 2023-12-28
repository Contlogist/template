package uc_user

import (
	"git.legchelife.ru/root/template/internal/repo/db/user"
	"git.legchelife.ru/root/template/pkg/models/context"
)

func (uc *UseCase) GetList(ctx context.Base, filter db_user.UserFilter) ([]db_user.User, error) {
	re, err := uc.repo.UserRepo.GetList(ctx, filter)
	if err != nil {
		return nil, err
	} else {
		return re, nil
	}
}
