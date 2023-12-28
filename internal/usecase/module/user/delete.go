package uc_user

import (
	"git.legchelife.ru/root/template/pkg/models/context"
)

func (uc *UseCase) Delete(ctx context.Base, id int) (bool, error) {
	re, err := uc.repo.UserRepo.Delete(ctx, id)
	if err != nil {
		return false, err
	} else {
		return re, nil
	}
}
