package uc_user

import "git.legchelife.ru/root/template/internal/repo"

type UserRepo struct {
	User  *User
	Param *Params
	Token *Token
}

func New(repo *repo.Repo) *UserRepo {
	return &UserRepo{
		User:  NewUser(repo),
		Param: NewParams(repo),
		Token: NewToken(repo),
	}
}
