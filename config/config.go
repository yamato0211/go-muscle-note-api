package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	UserName string
	Password string
	DBName   string
}

func LoadConfig() (config Config) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed load env")
	}

	return Config{
		Server: ServerConfig{Port: os.Getenv("PORT")},
		Database: DatabaseConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			UserName: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
		},
	}
}
