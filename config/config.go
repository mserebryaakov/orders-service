package config

import (
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Collection string `mapstructure:"collection"`
	AuthDB     string `mapstructure:"authdb"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Db     DbConfig     `mapstructure:"mongodb"`
	Server ServerConfig `mapstructure:"server"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()

	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config/.")

	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
