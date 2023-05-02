package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	MongoURI string
}

func GetConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	return &DBConfig{MongoURI: os.Getenv("MONGO_URI")}
}
