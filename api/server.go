package api

import (
	"context"
	"controller/api/utils"
	"controller/internal/messaging"
	"controller/internal/messaging/common"
	"github.com/fanap-infra/conf"

	"github.com/fanap-infra/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine    *gin.Engine
	Messaging messaging.Service
}

func (s *Server) Run() {
	addr := conf.GetString("addr", ":8090")
	log.Infov("Stream Service Runing", "addr", addr)
	utils.RunGin(s.Engine, addr)
}

func (s Server) KafkaInit() {
	err := s.Messaging.CreateTopic(context.Background(), []common.TopicSpecification{
		{
			Topic:             "Invoker1",
			NumPartitions:     1,
			ReplicationFactor: 1,
			ReplicaAssignment: nil,
			Config: map[string]string{
				"max.message.bytes":   "2000000",
				"segment.bytes":       "10000000",
				"retention.ms":        "5000",
				"segment.ms":          "10000",
				"delete.retention.ms": "10000",
			},
		},
	})
	if err != nil {
		log.Errorv("Create Topic", "topic", "Invoker1", "error", err)
	}
	err = s.Messaging.CreateTopic(context.Background(), []common.TopicSpecification{
		{
			Topic:             "Invoker2",
			NumPartitions:     1,
			ReplicationFactor: 1,
			ReplicaAssignment: nil,
			Config: map[string]string{
				"max.message.bytes":   "2000000",
				"segment.bytes":       "10000000",
				"retention.ms":        "5000",
				"segment.ms":          "10000",
				"delete.retention.ms": "10000",
			},
		},
	})
	if err != nil {
		log.Errorv("Create Topic", "topic", "Invoker2", "error", err)
	}

	err = s.Messaging.CreateTopic(context.Background(), []common.TopicSpecification{
		{
			Topic:             "Results",
			NumPartitions:     1,
			ReplicationFactor: 1,
			ReplicaAssignment: nil,
			Config: map[string]string{
				"max.message.bytes":   "2000000",
				"segment.bytes":       "10000000",
				"retention.ms":        "5000",
				"segment.ms":          "10000",
				"delete.retention.ms": "10000",
			},
		},
	})
	if err != nil {
		log.Errorv("Create Topic", "topic", "Results", "error", err)
	}
}
