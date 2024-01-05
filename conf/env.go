package conf

import (
	"fmt"
	"vorker/utils/secret"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	DBPath            string `env:"DB_PATH" env-default:"/workerd/db.sqlite"`
	WorkerdDir        string `env:"WORKERD_DIR" env-default:"/workerd"`
	DBType            string `env:"DB_TYPE" env-default:"sqlite"`
	WorkerLimit       int    `env:"WORKER_LIMIT" env-default:"10000"`
	WorkerdBinPath    string `env:"WORKERD_BIN_PATH" env-default:"/bin/workerd"`
	WorkerPort        int    `env:"WORKER_PORT" env-default:"8080"`
	APIPort           int    `env:"API_PORT" env-default:"8888"`
	ListenAddr        string `env:"LISTEN_ADDR" env-default:"0.0.0.0"`
	WorkerURLSuffix   string `env:"WORKER_URL_SUFFIX"`         // master required, e.g. .example.com. for worker show and route
	Scheme            string `env:"SCHEME" env-default:"http"` // http, https. for public frontend show
	CookieName        string `env:"COOKIE_NAME" env-default:"authorization"`
	CookieAge         int    `env:"COOKIE_AGE" env-default:"86400"` // second 86400 = 1 day
	CookieDomain      string `env:"COOKIE_DOMAIN"`                  // required, e.g. example.com
	EnableRegister    bool   `env:"ENABLE_REGISTER" env-default:"true"`
	AgentSecret       string `env:"AGENT_SECRET"` //	required, e.g. 123123123
	NodeName          string `env:"NODE_NAME" env-default:"default"`
	MasterEndpoint    string `env:"MASTER_ENDPOINT" env-default:"http://127.0.0.1:8888"` // needed for agent
	RunMode           string `env:"RUN_MODE" env-default:"master"`                       // master, agent
	TunnelEntryPort   int    `env:"TUNNEL_ENTRY_PORT" env-default:"10080"`
	TunnelHost        string `env:"TUNNEL_HOST" env-default:"127.0.0.1"` // for master usually 127.0.0.1, for agent usually master public ip
	TunnelAPIPort     int    `env:"TUNNEL_API_PORT" env-default:"18080"`
	DefaultWorkerHost string `env:"DEFAULT_WORKER_HOST" env-default:"localhost"`
	LitefsPrimaryPort int    `env:"LITEFS_PRIMARY_PORT" env-default:"20202"`
	TunnelUsername    string
	TunnelPassword    string
	TunnelToken       string
	NodeID            string
}

type JwtConfig struct {
	Secret     string `env:"JWT_SECRET" env-default:"secret"`
	ExpireTime int64  `env:"JWT_EXPIRETIME" env-default:"24"` // hour
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UID uint `json:"uid,omitempty"`
}

var (
	AppConfigInstance *AppConfig
	JwtConf           *JwtConfig
	RPCToken          string
)

func init() {
	var err error
	AppConfigInstance = &AppConfig{}
	JwtConf = &JwtConfig{}
	godotenv.Load()

	if err := cleanenv.ReadEnv(AppConfigInstance); err != nil {
		logrus.Panic(err)
	}
	if err := cleanenv.ReadEnv(JwtConf); err != nil {
		logrus.Panic(err)
	}

	RPCToken, err = secret.HashPassword(fmt.Sprintf("%s%s",
		AppConfigInstance.NodeName,
		AppConfigInstance.AgentSecret))
	AppConfigInstance.TunnelUsername = secret.MD5(AppConfigInstance.AgentSecret +
		AppConfigInstance.WorkerURLSuffix)
	AppConfigInstance.TunnelPassword = secret.MD5(AppConfigInstance.AgentSecret +
		AppConfigInstance.WorkerURLSuffix + AppConfigInstance.TunnelUsername)
	AppConfigInstance.TunnelToken = AppConfigInstance.TunnelPassword

	if err != nil {
		logrus.Panic(err)
	}
}

func IsMaster() bool {
	return AppConfigInstance.RunMode == "master"
}
