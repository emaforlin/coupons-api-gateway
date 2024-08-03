package internal

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Services []Service
}

type App struct {
	Debug   bool
	Version string
}
type Service struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
}

var debug bool

func LoadConfig() *Config {
	var services []Service
	if err := viper.UnmarshalKey("services", &services); err != nil {
		log.Fatal("unable to decode services", err)
	}

	config := Config{
		App: App{
			Debug:   debug,
			Version: viper.GetString("version"),
		},
		Services: services,
	}
	flag.Parse()
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
