package uc_user

import (
	"git.legchelife.ru/root/template/internal/repo/db/user"
	"git.legchelife.ru/root/template/pkg/models/context"
)

func (uc *UseCase) Put(ctx context.Base, user *db_user.User) (*db_user.User, error) {
	re, err := uc.repo.UserRepo.Put(ctx, user)
	if err != nil {
		return nil, err
	} else {
		return re, nil
	}
}
