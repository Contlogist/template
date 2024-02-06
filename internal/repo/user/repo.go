package user

import (
	"github.com/upper/db/v4"
)

type Repo struct {
	User  IUser
	Param IParam
	Token IToken
}

func New(db *db.Session) *Repo {
	return &Repo{
		&User{db: *db},
		&Param{db: *db},
		&Token{},
	}
}

func NewFake(
	UserRepo IUser,
	ParamRepo IParam,
	TokenRepo IToken,
) *Repo {
	return &Repo{
		UserRepo,
		ParamRepo,
		TokenRepo,
	}
}
