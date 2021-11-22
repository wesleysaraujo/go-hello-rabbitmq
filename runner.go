package main

import (
	"github.com/wesleysaraujo/go-hello-rabbitmq/consumer"
	"github.com/wesleysaraujo/go-hello-rabbitmq/publisher"
	"time"
)

func main() {
	publisher.Send()
	time.Sleep(2 * time.Second)
	consumer.Receive()
}
