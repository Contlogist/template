package gorm_db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dbURL string, models []interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// Распаковываем modules в список аргументов
	if err := db.AutoMigrate(models...); err != nil {
		logrus.Info("миграция не удалась")
		logrus.Error(err)
		return nil, err
	}
	logrus.Info("миграция прошла успешно")

	return db, nil
}
