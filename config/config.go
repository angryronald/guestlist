package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

var defaultConfig = map[string]string{
	// Common
	ENV: ENV_DEVELOPMENT,

	// Database
	DATABASE_CONNECTION_STRING: "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	DATABASE_HOST:              "localhost",
	DATABASE_PORT:              "3306",
	DATABASE_NAME:              "guestlist",
	DATABASE_USER:              "ronald",
	DATABASE_PASS:              "mysql_local",

	// Transport
	HTTP_ADDRESS: ":8001",
}

func GetEnv(KEY string) string {
	r := os.Getenv(KEY)

	if "" == r {
		if val, ok := defaultConfig[KEY]; ok {
			return val
		}

		return ""
	}

	return r
}

func GetFlavor() string {
	return GetEnv(ENV)
}

func GetWorkingDirectory() (string, error) {
	return os.Getwd()
}

func GetLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)

	return log
}
