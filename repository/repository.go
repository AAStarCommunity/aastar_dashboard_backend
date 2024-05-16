package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	dataBase *gorm.DB
	onlyOnce = sync.Once{}
)

func Init(dsn string) {
	onlyOnce.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logrus.Fatalf("Error when opening DB: %s\n", err)
		}
		dataBase = db
	})
}
