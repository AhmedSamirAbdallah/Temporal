package configs

import (
	"log"
	"os"
)

type Config struct {
	MongoURI string
	DBName   string
}

func LoadConfig() *Config {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
		log.Println("MONGO_URI not set, using default:", mongoURI)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "orders"
		log.Println("DB_NAME not set, using default:", dbName)
	}

	return &Config{
		MongoURI: mongoURI,
		DBName:   dbName,
	}
}
