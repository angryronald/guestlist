package config

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func DB() *gorm.DB {
	once.Do(func() {
		var err error

		dsn := getMySQLDSN()
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if nil != err {
			GetLogger().WithFields(logrus.Fields{
				"dsn":       dsn,
				"dialector": "postgres",
			}).Fatal("Failed to create DB Connection ", err.Error())
		}
	})

	return db
}

func getMySQLDSN() string {
	return fmt.Sprintf(GetEnv(DATABASE_CONNECTION_STRING), GetEnv(DATABASE_USER), GetEnv(DATABASE_PASS), GetEnv(DATABASE_HOST), GetEnv(DATABASE_PORT), GetEnv(DATABASE_NAME))
}
