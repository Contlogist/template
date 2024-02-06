package user

import (
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/upper"
	"github.com/upper/db/v4"
)

//go:generate mockgen -source=user.go -destination=./mock/user_mock.go -package=user_mock
type IUser interface {
	Get(ctx *context.Base, id int) (*User, error)
	GetList(ctx *context.Base, filter UserFilter) ([]User, error)
	Post(ctx *context.Base, u User) (*int, error)
	Put(ctx *context.Base, user *User) (bool, error)
	Delete(ctx *context.Base, id int) (bool, error)
}

type User struct {
	db           db.Session
	ID           int     `db:"id,omitempty" json:"id" swaggerignore:"true"`
	Name         string  `db:"name" json:"name" validate:"required"`
	Email        string  `db:"email" json:"email" validate:"required,email"`
	Password     string  `db:"password" json:"password" validate:"required"`
	CompanyID    int     `db:"company_id" json:"company_id" validate:"required"`
	RefreshToken *string `db:"refresh_token" json:"refresh_token"`
	//
	//Params *Params `db:"-" json:"-" swagger:"-"`
	//Tokens *Tokens `db:"-" json:"-" swagger:"-"`
}

// GetList получает список моделей User из базы данных, доступен фильтр UserFilter для поиска.
func (r *User) GetList(ctx *context.Base, filter UserFilter) ([]User, error) {
	ctx.SetTimeout(1)
	request, err := upper.DoRequest[[]User](ctx, func() ([]User, error) {
		users := make([]User, 0)
		sess := r.db.Collection("user")
		result := sess.Find(filter.Conditions())
		err := result.All(&users)
		if err != nil {
			return nil, err
		}

		return users, nil
	})
	if err != nil {
		return nil, err
	}
	return *request, nil
}

// Get получает модель User из базы данных по ID.
func (r *User) Get(ctx *context.Base, id int) (*User, error) {
	ctx.SetTimeout(3)
	request, err := upper.DoRequest[*User](ctx, func() (*User, error) {
		user := User{}
		err := r.db.Collection("user").Find(db.Cond{"id": id}).One(&user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	})
	if err != nil {
		return nil, err
	}
	return *request, nil
}

// Post создает новую модель User в базе данных и возвращает
func (r *User) Post(ctx *context.Base, user User) (*int, error) {
	ctx.SetTimeout(3)
	request, err := upper.DoRequest[*User](ctx, func() (*User, error) {
		err := r.db.Collection("user").InsertReturning(&user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	})
	if err != nil {
		return nil, err
	}
	return &(*request).ID, nil
}

// Put обновляет модель User в базе данных.
func (r *User) Put(ctx *context.Base, user *User) (bool, error) {
	ctx.SetTimeout(3)
	request, err := upper.DoRequest[bool](ctx, func() (bool, error) {
		err := r.db.Collection("user").Find(db.Cond{"id": user.ID}).Update(user)
		if err != nil {
			return false, err
		}
		return true, nil
	})
	if err != nil {
		return false, err
	}
	return *request, nil
}

// Delete удаляет модель User из базы данных.
func (r *User) Delete(ctx *context.Base, id int) (bool, error) {
	ctx.SetTimeout(3)
	request, err := upper.DoRequest[bool](ctx, func() (bool, error) {
		err := r.db.Collection("user").Find(db.Cond{"id": id}).Delete()
		if err != nil {
			return false, err
		}
		return true, nil
	})
	if err != nil {
		return false, err
	}
	return *request, nil
}
