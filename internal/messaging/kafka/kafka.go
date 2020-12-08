package kafka

import (
	"behnama/stream/services/messaging/common"
	"context"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fanap-infra/log"
)

type Service struct {
	log *log.Logger

	server string

	producer       *kafka.Producer
	producerLocker sync.Mutex
	producerCount  int

	readTimeout  time.Duration
	flushTimeout time.Duration
}

func New(server string, logger *log.Logger, readTimeout time.Duration, flushTimeout time.Duration) *Service {
	return &Service{
		log:          logger,
		server:       server,
		readTimeout:  readTimeout,
		flushTimeout: flushTimeout,
	}
}

func (s *Service) Close() {
	s.producerLocker.Lock()
	defer s.producerLocker.Unlock()
	s.producerCount = 0
	if s.producer != nil {
		s.log.Warn("Kafka force close producer")
		s.producer.Close()
		s.producer = nil
	}
}

func (s *Service) admin() (*kafka.AdminClient, error) {
	return kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": s.server,
	})
}

func (s *Service) CreateTopic(ctx context.Context, topics []common.TopicSpecification) error {
	admin, err := s.admin()
	if err != nil {
		return err
	}
	defer admin.Close()

	ts := make([]kafka.TopicSpecification, len(topics))

	for i, topic := range topics {
		ts[i].Topic = topic.Topic
		ts[i].NumPartitions = topic.NumPartitions
		ts[i].ReplicationFactor = topic.ReplicationFactor
		ts[i].ReplicaAssignment = topic.ReplicaAssignment
		ts[i].Config = topic.Config
	}

	_, err = admin.CreateTopics(ctx, ts, nil)
	return err
}

func (s *Service) Producer(topic string) (common.Producer, error) {
	s.producerLocker.Lock()
	defer s.producerLocker.Unlock()

	if s.producer == nil {
		s.log.Trace("Kafka create producer")
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": s.server,
		})
		if err != nil {
			return nil, err
		}
		s.producer = p
	}

	s.producerCount++
	return &producer{
		s:     s,
		topic: topic,
	}, nil
}

func (s *Service) closeProducer() {
	s.producerLocker.Lock()
	defer s.producerLocker.Unlock()
	s.producerCount--
	if s.producerCount == 0 && s.producer != nil {
		s.log.Trace("Kafka close producer")
		s.producer.Flush(int(s.flushTimeout.Milliseconds()))
		s.producer.Close()
		s.producer = nil
	}
}

func (s *Service) Consumer(groupID string) (common.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": s.server,
		"group.id":          groupID,
		// "auto.offset.reset": "end",
	})
	if err != nil {
		return nil, err
	}

	return &consumer{
		s:        s,
		Consumer: c,
	}, nil
}
