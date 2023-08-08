package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	KodoConfig Kodo
	DB         *gorm.DB
)

type Kodo struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

func init() {
	Config := readConfig()
	Dsn := getDSN(Config)
	DB = setupDatabase(Dsn)
	KodoConfig = loadKodoConfig(DB)
}

func readConfig() *viper.Viper {
	Config := viper.New()
	Config.SetConfigName("config")
	Config.SetConfigFile("config/config.toml")
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return Config
}

func getDSN(Config *viper.Viper) string {
	if Dsn, exists := os.LookupEnv("MYSQL_ADDR"); exists {
		hlog.Info("MYSQL_ADDR", Dsn)
		return Dsn
	}
	return Config.GetString("database.mysql_dsn")
}

func setupDatabase(Dsn string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return DB
}

func loadKodoConfig(DB *gorm.DB) Kodo {
	var kodoConfig Kodo
	DB.Table("system_config").Select("config_value").Where("config_key = ?", "kodo_access_key").Scan(&kodoConfig.AccessKey)
	DB.Table("system_config").Select("config_value").Where("config_key = ?", "kodo_secret_key").Scan(&kodoConfig.SecretKey)
	DB.Table("system_config").Select("config_value").Where("config_key = ?", "kodo_bucket_name").Scan(&kodoConfig.Bucket)
	DB.Table("system_config").Select("config_value").Where("config_key = ?", "kodo_domain").Scan(&kodoConfig.Domain)

	return kodoConfig
}
