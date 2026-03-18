package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 1
	go func() {
		for {
			msg := Message{i, "Hello from RabbitMQ"}
			atomic.AddInt64(&i, 1)
			c1 <- msg
		}
	}()
	go func() {
		for {
			msg := Message{i, "Hello from Kafka"}
			atomic.AddInt64(&i, 1)
			c1 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println(fmt.Sprintf("Received: ID: %d - MSGG: %s", msg.id, msg.Msg))

		case msg := <-c2:
			fmt.Println(fmt.Sprintf("Received: ID: %d - MSGG: %s", msg.id, msg.Msg))

		case <-time.After(time.Second * 3):
			fmt.Println("tiemout")

		}
	}
}
