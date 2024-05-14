package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DataBase *gorm.DB
)

func Init(dsn string) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Error when opening DB: %s\n", err)
	}
	DataBase = db
}
