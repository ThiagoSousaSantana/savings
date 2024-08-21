package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Config struct {
	API ApiConfig
	DB  DBConfig
}

type ApiConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewConfig(lc fx.Lifecycle, log *zap.Logger) (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	setDefaultValues()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Error("Error reading config file", zap.Error(err))
			return nil, err
		}
		log.Info("Config file not found", zap.String("file", viper.ConfigFileUsed()))
	} else {
		log.Info("Config loaded", zap.String("file", viper.ConfigFileUsed()))
	}

	fmt.Println(viper.GetString("api.port"))

	return &Config{
		API: ApiConfig{
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			Database: viper.GetString("db.database"),
		},
	}, nil
}

func setDefaultValues() {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "postgres")
	viper.SetDefault("db.database", "saving")
}
