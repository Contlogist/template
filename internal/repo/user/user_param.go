package user

import (
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/upper"
	"github.com/upper/db/v4"
)

//go:generate mockgen -source=user_param.go -destination=./mock/user_param_mock.go -package=user_mock
type IParam interface {
	GetList(ctx *context.Base, id int) ([]Param, error)
}

type Param struct {
	db     db.Session
	ID     int    `db:"id,omitempty" json:"id" swaggerignore:"true"`
	UserID int    `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
}

func (r *Param) GetList(ctx *context.Base, id int) ([]Param, error) {
	ctx.SetTimeout(3)
	request, err := upper.DoRequest[[]Param](ctx, func() ([]Param, error) {
		var userParams = make([]Param, 0)
		session := r.db.Collection("user_params")
		result := session.Find(db.Cond{"user_id": id})
		err := result.All(&userParams)
		if err != nil {
			return nil, err
		}
		return userParams, nil
	})
	if err != nil {
		return nil, err
	}
	return *request, nil
}
