package messaging

import (
	"fmt"
	"time"

	"github.com/fanap-infra/conf"
	"github.com/fanap-infra/log"
)

var configs = struct {
	Server       string
	ReadTimeout  time.Duration
	FlushTimeout time.Duration
}{
	Server:       "localhost",
	ReadTimeout:  time.Second * 2,
	FlushTimeout: time.Second * 3,
}

func LoadConfigs() {
	configs.ReadTimeout = conf.GetTimeDuration("messaging.timeout.read", configs.ReadTimeout)
	configs.FlushTimeout = conf.GetTimeDuration("messaging.timeout.flush", configs.FlushTimeout)
	configs.Server = conf.GetString("messaging.server", configs.Server)
	log.Infov("Messaging", "configs", fmt.Sprintf("%+v", configs))
}
