package config

import (
	"github.com/spf13/viper"
)

// Конфиг mongodb
type DbConfig struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Collection string `mapstructure:"collection"`
	AuthDB     string `mapstructure:"authdb"`
}

// Конфиг сервера
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// Конфиг приложения
type Config struct {
	Db     DbConfig     `mapstructure:"mongodb"`
	Server ServerConfig `mapstructure:"server"`
}

var vp *viper.Viper

// Загрузка конфига
func LoadConfig() (Config, error) {
	// Создание viper объекта конфига
	vp = viper.New()

	// Создание объекта конфига
	var config Config

	// Параметры конфига
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("config")

	// Чтение из config/config.json
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	// Десериализация json конфига
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
