package rp_company

import (
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/upper"
	"github.com/upper/db/v4"
)

func New(db *db.Session) *Company {
	return &Company{db: *db}
}

// Company model information.
// @Description Company model information.
type Company struct {
	db          db.Session
	ID          int    `db:"id,omitempty" json:"id" swaggerignore:"true"`
	Name        string `db:"name" json:"name" validate:"required"`
	Description string `db:"description" json:"description" validate:"required"`
	CreatedAt   string `db:"created_at" json:"created_at" swaggerignore:"true"`
}

func (r *Company) Get(ctx *context.Base, id int) (*Company, error) {
	ctx.SetTimeout(3)

	request, err := upper.DoRequest[*Company](ctx, func() (*Company, error) {
		company := Company{}
		err := r.db.Collection("user").Find(db.Cond{"id": id}).One(&company)
		if err != nil {
			return nil, err
		}
		return &company, nil
	})
	if err != nil {
		return nil, err
	}
	return *request, nil
}
