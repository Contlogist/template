package uc_user

import (
	"git.legchelife.ru/root/template/internal/repo/db/user"
	"git.legchelife.ru/root/template/pkg/models/context"
)

func (uc *UseCase) Post(ctx context.Base, u db_user.User) (*db_user.User, error) {
	re, err := uc.repo.UserRepo.Post(ctx, u)
	if err != nil {
		return nil, err
	} else {
		return re, nil
	}
}
