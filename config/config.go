package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
	"os"
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
	KodoConfig.AccessKey = os.Getenv("KODO_ACCESS_KEY")
	KodoConfig.SecretKey = os.Getenv("KODO_SECRET_KEY")
	//KodoConfig.Bucket = os.Getenv("KODO_BUCKET")
	//KodoConfig.Domain = os.Getenv("KODO_DOMAIN")
	if KodoConfig.AccessKey != "" && KodoConfig.SecretKey != "" {
		hlog.Info("kodo config success")
	} else {
		hlog.Info("kodo config failed")
	}
	KodoConfig.Bucket = Config.GetString("kodo.bucket")
	KodoConfig.Domain = Config.GetString("kodo.domain")
}
