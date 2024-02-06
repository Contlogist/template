package user

import (
	"github.com/upper/db/v4"
)

type UserFilter struct {
	CompanyID int    `json:"company_id" db:"company_id" swaggerignore:"true"`
	Email     string `json:"email" db:"email" example:"r.abramov@contlogist.ru" format:"email" required:"false"`
	Password  string `json:"password" db:"password" example:"123456" format:"password" required:"false"`
}

// Conditions возвращает условия для фильтрации для upper запроса в базу данных.
func (r *UserFilter) Conditions() db.Cond {
	conditions := db.Cond{}

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
