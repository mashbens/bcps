package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Port   string `toml:"port" mapstructure:"port"`
		JWTKey string `toml:"jwtkey"`
	} `toml:"app"`
	Database struct {
		Driver  string `toml:"driver"`
		DB_Host string `toml:"db_host"`
		DB_Port string `toml:"db_port"`
		DB_User string `toml:"db_user"`
		DB_Pass string `toml:"db_pass"`
		DB_Name string `toml:"db_name"`
	} `toml:"database"`
	Mailjet struct {
		PublicKey  string `toml:"publicKey"`
		PrivateKey string `toml:"privatekey"`
		Email      string `toml:"email"`
	} `toml:"mailjet"`
	Midtrans struct {
		ClientKey string `toml:"clientKey"`
		ServerKey string `toml:"serverKey"`
	} `toml:"midtrans"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.App.Port = "8081"

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("No config file found, using default config")

		return &defaultConfig
	}
	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("Failed to parse config file, using default config")
		return &defaultConfig
	}
	return &finalConfig
}
