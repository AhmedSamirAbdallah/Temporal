package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func NewKafkaProducer(producer sarama.SyncProducer) *KafkaProducer {
	return &KafkaProducer{Producer: producer}
}

func (kp *KafkaProducer) Publish(ctx context.Context, topic string, message interface{}) error {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error serializing message: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msgBytes),
	}
	partition, offset, err := kp.Producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("KafkaProduce : Message sent to partition %d with offset %d\n", partition, offset)
	return nil
}
