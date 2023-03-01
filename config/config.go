package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ApiPort string
	DbHost  string
	DbUser  string
	DbPass  string
	DbName  string
	DbPort  string
)

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed load env")
	}
	ApiPort = os.Getenv("PORT")
	DbHost = os.Getenv("POSTGRES_HOST")
	DbUser = os.Getenv("POSTGRES_USER")
	DbPass = os.Getenv("POSTGRES_PASSWORD")
	DbName = os.Getenv("POSTGRES_DB")
	DbPort = os.Getenv("POSTGRES_PORT")
}
