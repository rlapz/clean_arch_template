package config

import (
	"time"

	"github.com/spf13/viper"
)

/*
 * Db
 */
type Db struct {
	Host     string
	Port     uint16
	Name     string
	User     string
	Password string

	ConnectionLifetimeMax time.Duration
	ConnectionsOpenMax    int
	ConnectionsIdleMax    int
	Args                  string
}

func (d *Db) validate() {
	if len(d.Host) == 0 {
		panic("config: \"database.host\" should not empty!")
	}

	if len(d.Name) == 0 {
		panic("config: \"database.name\" should not empty!")
	}

	if len(d.User) == 0 {
		panic("config: \"database.user\" should not empty!")
	}

	if len(d.Password) == 0 {
		panic("config: \"database.password\" should not empty!")
	}
}

/*
 * Http
 */
type Http struct {
	Host      string
	Port      uint16
	IsPrefork bool
}

func (h *Http) validate() {
	if len(h.Host) == 0 {
		panic("config: \"http.host\" should not empty!")
	}
}

/*
 * App
 */
type App struct {
	Name    string
	LogFile string
}

func (a *App) validate() {
	if len(a.Name) == 0 {
		panic("config: \"app.name\" should not empty!")
	}
}

/*
 * Config
 */
type Config struct {
	App  App
	Http Http
	Db   Db
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

	ret := Config{
		App: App{
			Name:    config.GetString("app.name"),
			LogFile: config.GetString("app.log_file"),
		},
		Http: Http{
			Host:      config.GetString("http.host"),
			Port:      config.GetUint16("http.port"),
			IsPrefork: config.GetBool("http.is_prefork"),
		},
		Db: Db{
			Host:                  config.GetString("database.host"),
			Port:                  config.GetUint16("database.port"),
			Name:                  config.GetString("database.name"),
			User:                  config.GetString("database.User"),
			Password:              config.GetString("database.password"),
			ConnectionLifetimeMax: config.GetDuration("database.connection_lifetime_max"),
			ConnectionsOpenMax:    config.GetInt("database.connections_open_max"),
			ConnectionsIdleMax:    config.GetInt("database.connections_idle_max"),
			Args:                  config.GetString("database.args"),
		},
	}

	// validation
	ret.App.validate()
	ret.Http.validate()
	ret.Db.validate()

	return &ret, nil
}
