package db_user

type UserParams struct {
	ID   int    `db:"id,omitempty" json:"id" primarykey:"true" autoincrement:"true" swaggerignore:"true"`
	Name string `db:"name" json:"name"`
}

func (r *UserParams) Put() *UserParams {
	r.Name = "1"

	return r
}
