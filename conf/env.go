package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	DBPath     string
	WorkerdDir string
	DBType     string
}

var (
	AppConfigInstance AppConfig
)

func init() {
	AppConfigInstance = AppConfig{}
	godotenv.Load()

	if dbPath := os.Getenv("DB_PATH"); dbPath != "" {
		AppConfigInstance.DBPath = dbPath
		logrus.Println("DB_PATH environment variable is set")
	} else {
		logrus.Panic("DB_PATH environment variable is not set")
	}

	if workerdDir := os.Getenv("WORKERD_DIR"); workerdDir != "" {
		AppConfigInstance.WorkerdDir = workerdDir
		logrus.Println("WORKERD_DIR environment variable is set")
	} else {
		logrus.Panic("WORKERD_DIR environment variable is not set")
	}

	if dbType := os.Getenv("DB_TYPE"); dbType != "" {
		AppConfigInstance.DBType = dbType
		logrus.Println("DB_TYPE environment variable is set")
	} else {
		logrus.Panic("DB_TYPE environment variable is not set")
	}
}
