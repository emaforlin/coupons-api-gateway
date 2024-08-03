package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      App                `mapstructure:"app"`
	Services map[string]Service `mapstructure:"services"`
}

type App struct {
	Debug   bool
	Version string `mapstructure:"version"`
}

type Service struct {
	Host string `mapstructure:"host"`
}

var debug bool

func LoadConfig() *Config {
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error al mapear la configuración: %s", err)
	}

	config.App.Debug = debug

	return &config
}

func InitViper(filename string) {
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		panic("error reading config file")
	}
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.Parse()
}
