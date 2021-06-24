package config

import (
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	GrpcPort     string `env:"APP_GRPC_PORT,required"`
	DBHost       string `env:"DB_HOST,required"`
	DBPort       int    `env:"DB_PORT,required"`
	DBName       string `env:"DB_NAME,required"`
	DBUser       string `env:"DB_USER,required"`
	DBPassword   string `env:"DB_PASS,required"`
	DBDriverName string `env:"DB_DRIVER_NAME,required"`
}

var defaultConfig = Config{
	GrpcPort:     ":8888",
	DBHost:       "localhost",
	DBPort:       3000,
	DBName:       "sapawarga",
	DBUser:       "root",
	DBPassword:   "password",
	DBDriverName: "mysql",
}

func Get() *Config {
	cfg := loadEnv(defaultConfig)

	return cfg
}

func loadEnv(cfg Config) *Config {
	s := reflect.ValueOf(&cfg).Elem()
	typeOfT := s.Type()
	for i := 0; i < typeOfT.NumField(); i++ {
		fieldName := typeOfT.Field(i).Name

		if os.Getenv(fieldName) != "" {
			if s.Field(i).Kind() == reflect.String {
				s.Field(i).SetString(os.Getenv(fieldName))
			} else if s.Field(i).Kind() == reflect.Int {
				val, err := strconv.ParseInt(os.Getenv(fieldName), 10, 64)
				if nil != err {
					panic(err)
				}

				s.Field(i).SetInt(val)
			}
		}
	}
	return &cfg
}
