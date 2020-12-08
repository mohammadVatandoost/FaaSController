package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type consumer struct {
	s *Service
	*kafka.Consumer
}

func (c *consumer) Subscribe(topic string) error {
	return c.Subscribes([]string{topic})
}

func (c *consumer) Subscribes(topics []string) error {
	return c.SubscribeTopics(topics, nil)
}

func (c *consumer) Read(cancel <-chan struct{}) (key []byte, value []byte, err error) {
	for {
		select {
		case <-cancel:
			return
		default:
			m, err := c.ReadMessage(c.s.readTimeout)
			if err != nil {
				if err.(kafka.Error).Code() != kafka.ErrTimedOut {
					return nil, nil, err
				}
			} else {
				return m.Key, m.Value, nil
			}
		}
	}
}

func (c *consumer) Close() {
	_ = c.Consumer.Close()
}
