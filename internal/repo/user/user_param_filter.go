package user

import "github.com/upper/db/v4"

type UserPramFilter struct {
	Name      string `json:"name" db:"name"`
	CompanyID int    `json:"company_id" db:"company_id" swaggerignore:"true"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

// Conditions возвращает условия для фильтрации для upper запроса в базу данных.
func (r *UserPramFilter) Conditions() db.Cond {
	conditions := db.Cond{}

	if r.Name != "" {
		conditions["name"] = r.Name
	}

	if r.CompanyID != 0 {
		conditions["company_id"] = r.CompanyID
	}

	if r.Email != "" {
		conditions["email"] = r.Email
	}

	if r.Password != "" {
		conditions["password"] = r.Password
	}

	return conditions
}
