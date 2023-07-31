package config

import (
	"github.com/spf13/viper"
)

var (
	Dsn string
)

// Init initializes the config.
func Init() {
	Config := viper.New()
	Config.SetConfigName("config")
	Config.SetConfigFile("./config.toml")
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Dsn = Config.GetString("database.mysql_dsn")
}
