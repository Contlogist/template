package rp_section

import (
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/upper"
	"github.com/upper/db/v4"
)

func New(db *db.Session) *Section {
	return &Section{
		db: *db,
	}
}

type Section struct {
	db   db.Session
	ID   int    `db:"id,omitempty" json:"id" swaggerignore:"true"`
	Name string `db:"name" json:"name" validate:"required"`
	URL  string `db:"url" json:"url" validate:"required"`
	Icon string `db:"icon" json:"icon"`
}

func (r *Section) GetList(ctx *context.Base) ([]Section, error) {
	ctx.SetTimeout(3)

	request, err := upper.DoRequest[[]Section](ctx, func() ([]Section, error) {
		section := make([]Section, 0)
		err := r.db.Collection("section").Find().One(&section)
		if err != nil {
			return nil, err
		}
		return section, nil
	})
	if err != nil {
		return nil, err
	}
	return *request, nil
}
