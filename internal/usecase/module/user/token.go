package uc_user

import (
	"errors"
	"git.legchelife.ru/root/template/internal/repo"
	"git.legchelife.ru/root/template/internal/repo/user"
	"git.legchelife.ru/root/template/pkg/models/context"
)

type Token struct {
	repo *repo.Repo
}

func NewToken(repo *repo.Repo) *Token {
	return &Token{repo: repo}
}

// Get - метод получения токена user.Tokens по email и password. (авторизация)
func (uc *Token) Get(ctx *context.Base, email, password string) (*user.Token, error) {
	ctx.SetTimeout(3)

	filter := user.UserFilter{
		Email:    email,
		Password: password,
	}

	users, err := uc.repo.UserRepo.User.GetList(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		err := errors.New("пользователь не найден или неверный пароль")
		return nil, err
	}

	if len(users) > 1 {
		err := errors.New("найдено более одного пользователя")
		return nil, err
	}

	token, err := uc.repo.UserRepo.Token.Get(ctx, users[0].ID, users[0].CompanyID)
	if err != nil {
		return nil, err
	}

	users[0].RefreshToken = &token.Refresh.Token

	_, err = uc.repo.UserRepo.User.Put(ctx, &users[0])
	if err != nil {
		return nil, err
	}

	return token, nil
}
