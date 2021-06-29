package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	APP_ENV         string `env:"APP_ENV,required"`
	APP_GRPC_PORT   string `env:"APP_GRPC_PORT,required"`
	DB_HOST         string `env:"DB_HOST,required"`
	DB_PORT         int    `env:"DB_PORT,required"`
	DB_NAME         string `env:"DB_NAME,required"`
	DB_USER         string `env:"DB_USER,required"`
	DB_PASS         string `env:"DB_PASS,required"`
	DB_DRIVER_NAME  string `env:"DB_DRIVER_NAME,required"`
	DEBUG           bool   `env:"DEBUG,required"`
	JWT_SIGNING_KEY string `env:"JWT_SIGNING_KEY,required"`
}

var defaultConfig = &Config{
	APP_ENV:         "local",
	APP_GRPC_PORT:   ":8888",
	DB_HOST:         "localhost",
	DB_PORT:         3000,
	DB_NAME:         "sapawarga",
	DB_USER:         "root",
	DB_PASS:         "password",
	DB_DRIVER_NAME:  "mysql",
	DEBUG:           true,
	JWT_SIGNING_KEY: "secret",
}

func Get() *Config {
	var err error
	appEnv := os.Getenv(`APP_ENV`)
	appGRPCPort := os.Getenv(`APP_GRPC_PORT`)
	debugString := os.Getenv(`APP_DEBUG`)
	secret := os.Getenv("JWT_SIGNING_KEY")
	debug := false

	if debugString == "true" {
		debug = true
	}

	dbHost := os.Getenv(`DB_HOST`)
	dbPort, _ := strconv.Atoi(os.Getenv(`DB_PORT`))
	dbUser := os.Getenv(`DB_USER`)
	dbPassword := os.Getenv(`DB_PASS`)
	dbName := os.Getenv(`DB_NAME`)
	driverName := os.Getenv(`DB_DRIVER_NAME`)

	if appEnv == "" || appGRPCPort == "" {
		err = fmt.Errorf("[CONFIG][Critical] Please check section APP on enviroment config")
		panic(err)
	}

	defaultConfig.APP_GRPC_PORT = appGRPCPort
	defaultConfig.DEBUG = debug
	defaultConfig.JWT_SIGNING_KEY = secret

	if dbHost == "" || dbPort == 0 || dbUser == "" || dbName == "" || driverName == "" {
		err = fmt.Errorf("[CONFIG][Critical] Please check section DB on environtment config")
		panic(err)
	}

	defaultConfig.DB_HOST = dbHost
	defaultConfig.DB_PORT = dbPort
	defaultConfig.DB_USER = dbUser
	defaultConfig.DB_PASS = dbPassword
	defaultConfig.DB_NAME = dbName
	defaultConfig.DB_DRIVER_NAME = driverName

	return defaultConfig
}
