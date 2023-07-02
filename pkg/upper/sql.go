package upper

import (
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"time"
)

// New -.
func New(dbURL string) (db.Session, error) {
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
