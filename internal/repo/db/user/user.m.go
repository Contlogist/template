package db_user

import (
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/models/rechan"
	"github.com/upper/db/v4"
)

type UserRepo struct {
	db db.Session
}

func New(db db.Session) *UserRepo {
	return &UserRepo{db}
}

type User struct {
	ID     int        `db:"id,omitempty" json:"id" primarykey:"true" autoincrement:"true" swaggerignore:"true"`
	Name   string     `db:"name" json:"name"`
	Params UserParams ``
}

func (r *UserRepo) GetList(ctx context.Base, filter UserFilter) ([]User, error) {
	ctx.SetTimeout(3)
	reChan := make(chan rechan.Base)
	defer close(reChan)

	go func() {
		c := rechan.Base{}

		var users []User

		sess := r.db.Collection("user")

		conditions := db.Cond{}

		if filter.Name != "" {
			conditions["name"] = filter.Name
		}

		if filter.CompanyID != 0 {
			conditions["company_id"] = filter.CompanyID
		}

		result := sess.Find(conditions)
		err := result.All(&users)
		if err != nil {
			c.SendError(reChan, "GetUser - Find: ", err)
			return
		}

		c.SendData(reChan, users)
	}()

	select {
	case <-ctx.Context.Done():
		return nil, ctx.Context.Err()
	case re := <-reChan:
		if re.Error != nil {
			return nil, re.Error
		} else {
			return re.Data.([]User), nil
		}
	}
}
func (r *UserRepo) Get(ctx context.Base, id int) (*User, error) {
	ctx.SetTimeout(3)
	reChan := make(chan rechan.Base)
	defer close(reChan)

	go func() {
		c := rechan.Base{}

		user := User{}
		err := r.db.Collection("user").Find("id", id).One(&user)
		if err != nil {
			c.SendError(reChan, "GetUser - Find: ", err)
			return
		}
		c.SendData(reChan, user)
	}()

	select {
	case <-ctx.Context.Done():
		return nil, ctx.Context.Err()
	case re := <-reChan:
		if re.Error != nil {
			return nil, re.Error
		} else {
			user := re.Data.(User)
			return &user, nil
		}
	}
}
func (r *UserRepo) Post(ctx context.Base, user User) (*User, error) {
	ctx.SetTimeout(3)
	reChan := make(chan rechan.Base)
	defer close(reChan)

	go func() {
		c := rechan.Base{}

		err := r.db.Collection("user").InsertReturning(&user)
		if err != nil {
			c.SendError(reChan, "PostUser - InsertReturning: ", err)
			return
		}
		c.SendData(reChan, user)
	}()

	select {
	case <-ctx.Context.Done():
		return nil, ctx.Context.Err()
	case re := <-reChan:
		if re.Error != nil {
			return nil, re.Error
		} else {
			user := re.Data.(User)
			return &user, nil
		}
	}
}
func (r *UserRepo) Put(ctx context.Base, user *User) (*User, error) {
	ctx.SetTimeout(3)
	reChan := make(chan rechan.Base)
	defer close(reChan)

	go func() {
		c := rechan.Base{}

		err := r.db.Collection("user").UpdateReturning(user)
		if err != nil {
			c.SendError(reChan, "PutUser - UpdateReturning: ", err)
			return
		}
		c.SendData(reChan, user)
	}()

	select {
	case <-ctx.Context.Done():
		return nil, ctx.Context.Err()
	case re := <-reChan:
		if re.Error != nil {
			return nil, re.Error
		} else {
			return re.Data.(*User), nil
		}
	}
}
func (r *UserRepo) Delete(ctx context.Base, id int) (bool, error) {
	ctx.SetTimeout(3)
	reChan := make(chan rechan.Base)
	defer close(reChan)

	go func() {
		c := rechan.Base{}

		err := r.db.Collection("user").Find(db.Cond{"id": id}).Delete()
		if err != nil {
			c.SendError(reChan, "GetUser - Find: ", err)
		}

		c.SendData(reChan, true)
	}()

	select {
	case <-ctx.Context.Done():
		return false, ctx.Context.Err()
	case re := <-reChan:
		if re.Error != nil {
			return false, re.Error
		} else {
			return re.Data.(bool), nil
		}
	}
}
