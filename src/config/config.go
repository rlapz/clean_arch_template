package config

import (
	"time"

	"github.com/spf13/viper"
)

type Db struct {
	Host     string
	Port     uint16
	Name     string
	User     string
	Password string

	ConnectionLifetimeMax time.Duration
	ConnectionsOpenMax    int
	ConnectionsIdleMax    int
}

type Http struct {
	Host      string
	Port      uint16
	IsPrefork bool
}

type Config struct {
	AppName string
	Http    Http
	Db      Db
}

func Load(isProduction bool) (*Config, error) {
	configName := "config.development"
	if isProduction {
		configName = "config.production"
	}

	config := viper.New()
	config.SetConfigName(configName)
	config.SetConfigType("json")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		AppName: viper.GetString("app.name"),
		Http: Http{
			Host:      config.GetString("http.host"),
			Port:      config.GetUint16("http.port"),
			IsPrefork: config.GetBool("http.is_prefork"),
		},
		Db: Db{
			Host:                  config.GetString("db.host"),
			Port:                  config.GetUint16("db.port"),
			Name:                  config.GetString("db.name"),
			User:                  config.GetString("db.User"),
			Password:              config.GetString("db.password"),
			ConnectionLifetimeMax: config.GetDuration("db.connection_lifetime_max"),
			ConnectionsOpenMax:    config.GetInt("db.connections_open_max"),
			ConnectionsIdleMax:    config.GetInt("db.connections_idle_max"),
		},
	}, nil
}
