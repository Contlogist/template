package uc_user

import (
	"git.legchelife.ru/root/template/internal/repo/db/user"
	"git.legchelife.ru/root/template/pkg/models/context"
)

func (uc *UseCase) Get(ctx context.Base, id int) (*db_user.User, error) {
	re, err := uc.repo.UserRepo.Get(ctx, id)

	userParams := db_user.UserParams{
		ID:   1,
		Name: "123",
	}

	user := db_user.User{
		ID:     1,
		Name:   "123",
		Params: userParams,
	}

	user.Params.Put()

	if err != nil {
		return nil, err
	} else {
		return re, nil
	}
}
