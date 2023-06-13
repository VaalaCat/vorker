package conf

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	DBPath          string `env:"DB_PATH"`
	WorkerdDir      string `env:"WORKERD_DIR"`
	DBType          string `env:"DB_TYPE" env-default:"sqlite"`
	WorkerLimit     int    `env:"WORKER_LIMIT" env-default:"1000"`
	WorkerdBinPath  string `env:"WORKERD_BIN_PATH"`
	WorkerPort      int    `env:"WORKER_PORT" env-default:"8080"`
	APIPort         int    `env:"API_PORT" env-default:"8888"`
	ListenAddr      string `env:"LISTEN_ADDR" env-default:"0.0.0.0"`
	WorkerURLSuffix string `env:"WORKER_URL_SUFFIX"`
	Scheme          string `env:"SCHEME" env-default:"http"`
	CookieName      string `env:"COOKIE_NAME" env-default:"authorization"`
	CookieAge       int    `env:"COOKIE_AGE" env-default:"86400"` // sec
	CookieDomain    string `env:"COOKIE_DOMAIN"`
	EnableRegister  bool   `env:"ENABLE_REGISTER" env-default:"true"`
	TunnelPort      int    `env:"TUNNEL_PORT" env-default:"6666"`
}

type JwtConfig struct {
	Secret     string `env:"JWT_SECRET"`
	ExpireTime int64  `env:"JWT_EXPIRETIME"` // hour
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UID uint `json:"uid,omitempty"`
}

var (
	AppConfigInstance *AppConfig
	JwtConf           *JwtConfig
)

func init() {
	AppConfigInstance = &AppConfig{}
	JwtConf = &JwtConfig{}
	godotenv.Load()

	if err := cleanenv.ReadEnv(AppConfigInstance); err != nil {
		logrus.Panic(err)
	}
	if err := cleanenv.ReadEnv(JwtConf); err != nil {
		logrus.Panic(err)
	}
}
