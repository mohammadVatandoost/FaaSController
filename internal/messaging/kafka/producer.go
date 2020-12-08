package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type producer struct {
	s     *Service
	topic string
}

func (p *producer) Publish(payload []byte) error {
	return p.s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}, nil)
}

func (p *producer) Close() {
	p.s.closeProducer()
}
