package main

import (
	"controller/api"
	"controller/api/utils"
	"controller/internal/messaging"
	limits "github.com/gin-contrib/size"
)

func main() {
	var server api.Server
	server.Engine = utils.NewGinEngine(true)
	server.Engine.Use(limits.RequestSizeLimiter(100000))
	server.Messaging = messaging.New()
	defer server.Messaging.Close()
	server.KafkaInit()
	server.Routes()
	server.Run()
}
