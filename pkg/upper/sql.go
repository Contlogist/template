package upper

import (
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"time"
)

// NewSQL -.
func NewSQL(dbURL string) (db.Session, error) {
	settings, _ := mysql.ParseURL(dbURL)
	sess, err := mysql.Open(settings)
	sess.SetConnMaxLifetime(time.Minute * 4)
	if err != nil {
		return nil, err
	}
	if err := sess.Ping(); err != nil {
		return nil, err
	}
	return sess, nil
}

// NewPostgres -.
func NewPostgres(dbURL string) (db.Session, error) {
	settings, _ := postgresql.ParseURL(dbURL)
	sess, err := postgresql.Open(settings)
	sess.SetConnMaxLifetime(time.Minute * 4)
	if err != nil {
		return nil, err
	}
	if err := sess.Ping(); err != nil {
		return nil, err
	}
	return sess, nil
}
