package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort     string
	MongoURI       string
	DBName         string
	CollectionName string
	KafkaBrokers   []string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":6000"
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

	kafkaBrokersEnv := os.Getenv("KAFKA_BROKERS")
	var kafkaBrokers []string
	if kafkaBrokersEnv == "" {
		kafkaBrokers = []string{"localhost:9092"}
		log.Println("KAFKA_BROKERS not set, using default:", kafkaBrokers)
	} else {
		kafkaBrokers = strings.Split(kafkaBrokersEnv, ",")
	}

	return &Config{
		ServerPort:     serverPort,
		MongoURI:       mongoURI,
		DBName:         dbName,
		CollectionName: collectionName,
		KafkaBrokers:   kafkaBrokers,
	}
}
