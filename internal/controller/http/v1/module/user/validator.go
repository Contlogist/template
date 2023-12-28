package rt_user

import (
	"errors"
	db_user "git.legchelife.ru/root/template/internal/repo/db/user"
)

func (r *Router) vGet(id int) error {
	if id == 0 {
		return errors.New("id is empty")
	}

	return nil
}

func (r *Router) vGetList(filter db_user.UserFilter) error {
	return nil
}

func (r *Router) vPost(user db_user.User) error {
	if user.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}

func (r *Router) vPut(user db_user.User) error {
	if user.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}

func (r *Router) vDelete(id int) error {
	if id == 0 {
		return errors.New("id is empty")
	}

	return nil
}
