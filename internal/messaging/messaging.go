package messaging

import (

	"context"
	"controller/internal/messaging/common"
	"controller/internal/messaging/kafka"
	"github.com/fanap-infra/log"
)

type Service interface {
	CreateTopic(ctx context.Context, topics []common.TopicSpecification) error

	Producer(topic string) (common.Producer, error)
	Consumer(groupID string) (common.Consumer, error)

	Close()
}

func New() Service {
	return kafka.New(configs.Server, log.GetScope("Messaging"), configs.ReadTimeout, configs.FlushTimeout)
}
