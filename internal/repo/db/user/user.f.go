package db_user

type UserFilter struct {
	Name      string `json:"name" db:"name"`
	CompanyID int    `json:"company_id" db:"company_id"`
}
