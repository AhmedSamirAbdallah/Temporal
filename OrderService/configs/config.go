package configs

import (
	"log"
	"os"
)

type Config struct {
	ServerPort     string
	MongoURI       string
	DBName         string
	CollectionName string
}

func LoadConfig() *Config {

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "6000"
		log.Println("SERVER_PORT not set, using default:", serverPort)
	}

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

	collectionName := os.Getenv("COLLERCTION_NAME")
	if collectionName == "" {
		collectionName = "order"
		log.Println("COLLERCTION_NAME not set, using default:", collectionName)
	}

	return &Config{
		ServerPort:     serverPort,
		MongoURI:       mongoURI,
		DBName:         dbName,
		CollectionName: collectionName,
	}
}
