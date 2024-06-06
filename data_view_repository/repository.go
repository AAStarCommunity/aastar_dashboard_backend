package data_view_repository

import (
	"aastar_dashboard_back/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	dataVeiewDB *gorm.DB
	onlyOnce    = sync.Once{}
)

func Init() {
	newLogger := logger.New(
		logrus.New(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	onlyOnce.Do(func() {

		dsn := config.GetDataViewDsn()
		if dsn == "" {
			logrus.Fatalf("DSN is empty")
		}
		logrus.Infof("DSN : %s", dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			logrus.Fatalf("Error when opening DB: %s\n", err)
		}
		dataVeiewDB = db
	})
}
