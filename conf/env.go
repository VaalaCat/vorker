package conf

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	DBPath         string
	WorkerdDir     string
	DBType         string
	WorkerLimit    int `json:"WorkerLimit"`
	WorkerdBinPath string
	ProxyPort      int
	APIPort        int
	ListenAddr     string
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

	if workerLimit := os.Getenv("WORKER_LIMIT"); workerLimit != "" {
		WorkerLimitNum, err := strconv.Atoi(workerLimit)
		if err != nil {
			logrus.Panic("WORKER_LIMIT environment variable is not a number")
		}
		if WorkerLimitNum < 1 {
			logrus.Panic("WORKER_LIMIT environment variable is less than 1")
		}
		AppConfigInstance.WorkerLimit = WorkerLimitNum
		logrus.Println("WORKER_LIMIT environment variable is set")
	} else {
		logrus.Panic("WORKER_LIMIT environment variable is not set")
	}

	if AppConfigInstance.WorkerLimit > 65535 {
		logrus.Panic("WORKER_LIMIT is too large")
	}

	if workerdBinPath := os.Getenv("WORKERD_BIN_PATH"); workerdBinPath != "" {
		AppConfigInstance.WorkerdBinPath = workerdBinPath
		logrus.Println("WORKERD_BIN_PATH environment variable is set")
	} else {
		logrus.Panic("WORKERD_BIN_PATH environment variable is not set")
	}

	if proxyPort := os.Getenv("PROXY_PORT"); proxyPort != "" {
		proxyPortNum, err := strconv.Atoi(proxyPort)
		if err != nil {
			logrus.Panic("PROXY_PORT environment variable is not a number")
		}
		if proxyPortNum < 1 || proxyPortNum > 65535 {
			logrus.Panic("PROXY_PORT environment variable is not valid")
		}
		AppConfigInstance.ProxyPort = proxyPortNum
		logrus.Println("PROXY_PORT environment variable is set")
	} else {
		logrus.Panic("PROXY_PORT environment variable is not set")
	}

	if apiPort := os.Getenv("API_PORT"); apiPort != "" {
		apiPortNum, err := strconv.Atoi(apiPort)
		if err != nil {
			logrus.Panic("API_PORT environment variable is not a number")
		}
		if apiPortNum < 1 || apiPortNum > 65535 {
			logrus.Panic("API_PORT environment variable is not valid")
		}
		AppConfigInstance.APIPort = apiPortNum
		logrus.Println("API_PORT environment variable is set")
	} else {
		logrus.Panic("API_PORT environment variable is not set")
	}

	if listenAddr := os.Getenv("LISTEN_ADDR"); listenAddr != "" {
		AppConfigInstance.ListenAddr = listenAddr
		logrus.Println("LISTEN_ADDR environment variable is set")
	} else {
		logrus.Panic("LISTEN_ADDR environment variable is not set")
	}
}
