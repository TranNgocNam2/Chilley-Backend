package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	CorsDev  string
	CorsProd string
}

func LoadAllAppConfig() (config *Config, err error) {
	fmt.Println(os.Getenv("APP_HOST"))
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal("Failed to parse port")
	}
	config = &Config{
		Host:     os.Getenv("APP_HOST"),
		Port:     port,
		CorsDev:  os.Getenv("CORS_DEV"),
		CorsProd: os.Getenv("CORS_PROD"),
	}
	return config, nil
}
