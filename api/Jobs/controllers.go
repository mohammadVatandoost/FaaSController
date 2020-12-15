package Jobs

import (
	invokerMessage "controller/assets/protobuf"
	"controller/internal/messaging"
	"github.com/fanap-infra/log"
	"github.com/golang/protobuf/proto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Controller ...
type Controller struct {
	Messaging messaging.Service
}

type CacheItem struct{
	ID int
	CreatedTime time.Time
	LastTimeUsed time.Time
	Result []byte
}

type JobReq struct {
	ID        int32   `json:"id"`
	Variables []string `json:"variables"`
}

func (jobCon Controller) jobHandler(c *gin.Context) {
	var jobReq JobReq
	err := c.BindJSON(&jobReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	jReq := &invokerMessage.JobRequest{
		Id: jobReq.ID,
	}
	data, err := proto.Marshal(jReq)
	if err != nil {
		log.Errorv("jobHandler marshaling ", "error", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	producer, err :=jobCon.Messaging.Producer("Invoker1")
	if err != nil {
		log.Errorv("jobHandler Messaging.Producer ", "error", err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	if err = producer.Publish(data); err != nil {
		log.Errorv("jobHandler messaging publish", "topic", "Invoker1",  "error", err)
		c.JSON(http.StatusBadRequest, "")
	}

	go func() {
		consumer, err := jobCon.Messaging.Consumer("Results")
		if err != nil {
			log.Errorv("jobHandler messaging consumer", "error", err)
			c.JSON(http.StatusBadRequest, "")
			return
		}

		if err = consumer.Subscribe("Results"); err != nil {
			log.Errorv("jobHandler messaging subscribe", "topic", "Results", "error", err)
			c.JSON(http.StatusBadRequest, "")
			return
		}
        stop := make(chan struct{})
		for {
			select {
			case <-stop:
				c.JSON(http.StatusBadRequest, "")
				return
			default:
				if key, value, e := consumer.Read(stop); e != nil {
					log.Errorv("jobHandler messaging read", "error", e)
				} else if key != nil && value != nil {
					log.Info("jobHandler messaging read  Ok***")
					c.JSON(http.StatusOK, "")
					return
				} else {
					log.Warnv("jobHandler messaging Not read", "key", key, "value", value)
				}
			}

		}
	}()


}
