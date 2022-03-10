package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("My_NSQ_Topic", "My_NSQ_Channel", decodeConfig)
	if err != nil {
		log.Panicln("Could not create consumer")
	}
	//c.MaxInFlight defaults to 1

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("NSQ message received: %v", string(message.Body))
		return nil
	}))

	err = c.ConnectToNSQDs([]string{"localhost:4150"})
	if err != nil {
		log.Panicln("Could not connect to nsq daemon")
	}
	log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
	wg.Wait()
}
