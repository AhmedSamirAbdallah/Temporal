package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func SetupKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	return config
}

func ConnectKafka(brokers []string, config *sarama.Config) (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Printf("Failed to connect to Kafka: %v", err)
		return nil, err
	}
	log.Println("Connected to Kafka successfully")
	return producer, nil
}
