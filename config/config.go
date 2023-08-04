package config

import (
	"github.com/spf13/viper"
)

var (
	Dsn        string
	KodoConfig Kodo
)

type Kodo struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

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
	KodoConfig.AccessKey = Config.GetString("kodo.access_key")
	KodoConfig.SecretKey = Config.GetString("kodo.secret_key")
	KodoConfig.Bucket = Config.GetString("kodo.bucket")
	KodoConfig.Domain = Config.GetString("kodo.domain")
}
