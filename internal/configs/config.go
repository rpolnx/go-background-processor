package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type AppProfile string

const (
	LocalProfile AppProfile = "LOCAL"
	DevProfile   AppProfile = "DEV"
	TestProfile  AppProfile = "TEST"
	ProdProfile  AppProfile = "PROD"
)

type AppConfig struct {
	Host    string     `env:"HOST"`
	Port    int        `env:"PORT"`
	Profile AppProfile `env:"PROFILE"`
}

var GlobalAppConfig *AppConfig = &AppConfig{}

func InitEnvVariables() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	err = env.Parse(GlobalAppConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to parse environment variables: %v", err)
	}

	return GlobalAppConfig, nil
}
