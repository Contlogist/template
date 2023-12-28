package repo_db

import (
	"git.legchelife.ru/root/template/internal/repo/db/user"
	"github.com/upper/db/v4"
)

type RepoDB struct {
	db       *db.Session
	UserRepo db_user.UserRepo
}

func New(db *db.Session) *RepoDB {

	userRepo := db_user.New(*db)

	return &RepoDB{
		db,
		*userRepo,
	}
}
