package configs

import (
	"errors"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type apiConfig struct {
	Env  string
	Port int
}

type dbConfig struct {
	Postgres dbConfigParams
}

type dbConfigParams struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SslMode  string
}

var (
	APIConfig apiConfig
	DBConfig  dbConfig
)

func init() {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("absolute path for the configs folder could not be retrieved")
	}

	rootDir := path.Join(path.Dir(f), "../..")
	viper.AddConfigPath(rootDir + "/")

	switch env := os.Getenv("ENV"); env {
	case "PRODUCTION":
		viper.SetConfigName("config.production")
	default:
		viper.SetConfigName("config.development")
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		var fnfErr viper.ConfigFileNotFoundError
		if errors.As(err, &fnfErr) {
			log.Fatalf("config file not found: %s", fnfErr)
		}

		log.Fatalf("error reading config file: %s", err)
	}

	APIConfig = apiConfig{
		Env:  viper.GetString("api.environment"),
		Port: viper.GetInt("api.port"),
	}

	DBConfig = dbConfig{
		Postgres: dbConfigParams{
			Host:     viper.GetString("database.postgres.host"),
			Port:     viper.GetInt("database.postgres.port"),
			User:     viper.GetString("database.postgres.user"),
			Password: viper.GetString("database.postgres.password"),
			Name:     viper.GetString("database.postgres.name"),
			SslMode:  viper.GetString("database.postgres.sslmode"),
		},
	}

	log.Println("environment is ready")
}

func (a apiConfig) IsDevelopment() bool {
	if a.Env == "DEVELOPMENT" {
		return true
	} else {
		return false
	}
}

func (a apiConfig) IsProduction() bool {
	if a.Env == "PRODUCTION" {
		return true
	} else {
		return false
	}
}
