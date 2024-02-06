package uc_user

import (
	"git.legchelife.ru/root/template/internal/repo"
	"git.legchelife.ru/root/template/internal/repo/user"
	"git.legchelife.ru/root/template/pkg/models/context"
)

type User struct {
	repo *repo.Repo
	//Params *Params
	//Token  *Token
}

func NewUser(repo *repo.Repo) *User {
	return &User{
		repo: repo,
	}
}

// Get получает модель User из репозитория по ID.
func (uc *User) Get(ctx *context.Base, id int) (*user.User, error) {
	//ctx.SetTimeout(3)
	re, err := uc.repo.UserRepo.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// GetList получает список моделей User из репозитория, доступен фильтр user.UserFilter для поиска.
func (uc *User) GetList(ctx *context.Base, filter user.UserFilter) ([]user.User, error) {
	//ctx.SetTimeout(3)
	re, err := uc.repo.UserRepo.User.GetList(ctx, filter)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// Post создает новую модель User в репозитории.
func (uc *User) Post(ctx *context.Base, u user.User) (*int, error) {
	//ctx.SetTimeout(3)
	re, err := uc.repo.UserRepo.User.Post(ctx, u)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// Put обновляет модель User в репозитории.
func (uc *User) Put(ctx *context.Base, user *user.User) (bool, error) {
	//ctx.SetTimeout(3)
	re, err := uc.repo.UserRepo.User.Put(ctx, user)
	if err != nil {
		return false, err
	}
	return re, nil
}

// Delete удаляет модель User из репозитория по ID.
func (uc *User) Delete(ctx *context.Base, id int) (bool, error) {
	//ctx.SetTimeout(3)
	re, err := uc.repo.UserRepo.User.Delete(ctx, id)
	if err != nil {
		return false, err
	}
	return re, nil
}
