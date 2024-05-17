package repository

import (
	"aastar_dashboard_back/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	dataBase *gorm.DB
	onlyOnce = sync.Once{}
)

func Init() {

	onlyOnce.Do(func() {
		dsn := config.GetDsn()
		logrus.Infof("DSN : %s", dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logrus.Fatalf("Error when opening DB: %s\n", err)
		}
		dataBase = db
	})
}
