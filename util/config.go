package util

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

type AbstractConfig interface {
	Load() *GenericConfig
}

func LoadGeneralConfig(confName string) (*GenericConfig, error) {
	var genericConfig GenericConfig

	viperConfig, err := LoadConfig(".")
	if err != nil {
		return nil, err
	}
	genericConfig.AppName = viperConfig.AppName

	return &genericConfig, nil
}

type GenericConfig struct {
	DBDriver      string `json:"dbd_river"`
	DBSource      string `json:"dbd_ource"`
	ServerAddress string `json:"server_address"`

	AppName string `json:"app_name"`
	AppEnv  string `json:"app_env"`
	AppUrl  string `json:"app_url"`

	AppHost          string `json:"app_host"`
	AppFrontendHost  string `json:"app_frontendHost"`
	AppGitName       string `json:"app_gitName"`
	AppStorageFolder string `json:"app_storageFolder"`

	DbDriver   string `json:"db_driver"`
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbDatabase string `json:"db_database"`
	DbUsername string `json:"db_username"`
	DbPassword string `json:"db_password"`

	Oauth2Client          string `json:"oauth2_client"`
	Oauth2Secret          string `json:"oauth2_secret"`
	FrontendRedirectToken string `json:"frontend_redirect_token"`

	TokenSymmetricKey   string        `json:"token_symmetric_key"`
	AccessTokenDuration time.Duration `json:"time"`

	SMTPHost  string `json:"smtp_host"`
	SMTPPort  int    `json:"smtp_port"`
	SMTPUser  string `json:"smtp_user"`
	SMTPPass  string `json:"smtp_pass"`
	EmailFrom string `json:"email_from"`
}

type Config struct {
	AppName string `mapstructure:"APP_NAME"`

	TwilioFrom      string `mapstructure:"TWILIO_FROM"`
	TwilioTo        string `mapstructure:"TWILIO_TO"`
	TwilioSID       string `mapstructure:"TWILIO_SID"`
	TwilioAuthToken string `mapstructure:"TWILIO_AUTH_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	// env := os.Getenv("ENV")
	// if env == "" {
	// 	env = ".env"
	// }

	// // Carregue as variáveis de ambiente apropriadas do arquivo .env correspondente
	// err = godotenv.Load(fmt.Sprintf(".env.%s", env))
	// if err != nil {
	// 	panic(fmt.Errorf("Erro ao carregar variáveis de ambiente: %s", err))
	// }

	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Erro ao carregar arquivo de configuração: %s", err))
	}

	err = viper.Unmarshal(&config)
	return
}

func PrintMyPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func CheckPassword(passwd1 string, passwd2 string) error {

	return nil
}
