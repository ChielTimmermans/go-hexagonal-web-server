package main

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name string
		URL  string
		Port int
	}
	Client struct {
		ReadTimeout     time.Duration
		ReadBufferSize  int
		WriteTimeout    time.Duration
		WriteBufferSize int
	}
	CORS     *ConfigCORS
	Language *ConfigLanguage
	Mail     *ConfigMail
	Router   *ConfigRouter

	DBPostgres *ConfigStoragePostgres

	URL struct {
		App       string
		Dashboard string
	}
}

type ConfigRouter struct {
	RedirectTrailingSlash  bool
	RedirectFixedPath      bool
	HandleMethodNotAllowed bool
	HandleOPTIONS          bool
}

type ConfigCORS struct {
	AllowedOrigins   []string
	AllowedHeaders   []string
	AllowCredentials bool
	Debug            bool
	AllowMaxAge      int
}

type ConfigMail struct {
	Identity string
	Username string
	Password string
	Host     string
	Port     int
	From     string
}

type ConfigLanguage struct {
	AvailableLanguages []string
	DefaultLanguage    string
}

type ConfigStoragePostgres struct {
	Hostname string
	Port     int
	User     string
	Password string
	Database string
}

func initConfig() (config *Config) {
	log.Println("Init config")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	config = &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic("unable to decode into config struct")
	}
	return
}
