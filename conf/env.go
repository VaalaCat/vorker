package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
		log.Println("DB_PATH environment variable is set")
	} else {
		log.Panic("DB_PATH environment variable is not set")
	}

	if workerdDir := os.Getenv("WORKERD_DIR"); workerdDir != "" {
		AppConfigInstance.WorkerdDir = workerdDir
		log.Println("WORKERD_DIR environment variable is set")
	} else {
		log.Panic("WORKERD_DIR environment variable is not set")
	}

	if dbType := os.Getenv("DB_TYPE"); dbType != "" {
		AppConfigInstance.DBType = dbType
		log.Println("DB_TYPE environment variable is set")
	} else {
		log.Panic("DB_TYPE environment variable is not set")
	}
}
