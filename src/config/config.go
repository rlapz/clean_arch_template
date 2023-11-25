package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string
	Port     uint16
	Name     string
	Password string
}

type HttpConfig struct {
	Host      string
	Port      uint16
	IsPrefork bool
}

type Config struct {
	AppName  string
	Http     HttpConfig
	Db       DbConfig
	Log      *logrus.Logger
	Validate *validator.Validate
}

func Load(isProduction bool) (*Config, error) {
	configName := "config.development"
	if isProduction {
		configName = "config.production"
	}

	viperConfig := viper.New()
	viperConfig.SetConfigName(configName)
	viperConfig.SetConfigType("json")
	viperConfig.AddConfigPath(".")

	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		AppName: viper.GetString("app.name"),
		Http: HttpConfig{
			Host:      viperConfig.GetString("http.host"),
			Port:      viperConfig.GetUint16("http.port"),
			IsPrefork: viperConfig.GetBool("http.is_prefork"),
		},
		Db: DbConfig{
			Host:     viperConfig.GetString("db.host"),
			Port:     viperConfig.GetUint16("db.port"),
			Name:     viperConfig.GetString("db.name"),
			Password: viperConfig.GetString("db.password"),
		},
		Log:      logrus.New(),
		Validate: validator.New(),
	}, nil
}
