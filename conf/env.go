package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	DBPath             string `env:"DB_PATH"`
	WorkerdDir         string `env:"WORKERD_DIR"`
	DBType             string `env:"DB_TYPE"`
	WorkerLimit        int    `env:"WORKER_LIMIT"`
	WorkerdBinPath     string `env:"WORKERD_BIN_PATH"`
	WorkerPort         int    `env:"WORKER_PORT"`
	APIPort            int    `env:"API_PORT"`
	ListenAddr         string `env:"LISTEN_ADDR"`
	WorkerURLSuffix string `env:"WORKER_URL_SUFFIX"`
	Scheme             string `env:"SCHEME"`
}

var (
	AppConfigInstance AppConfig
)

func init() {
	AppConfigInstance = AppConfig{}
	godotenv.Load()
	err := cleanenv.ReadEnv(&AppConfigInstance)
	if err != nil {
		logrus.Panic(err)
	}
}
