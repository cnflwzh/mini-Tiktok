package config

import (
	"github.com/spf13/viper"
)

var (
	Dsn string
)

// init initializes the config.
func init() {
	Config := viper.New()
	Config.SetConfigName("config")
	Config.SetConfigFile("config/config.toml")
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Dsn = Config.GetString("database.mysql_dsn")
}
