package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("localhost:4150", config)
	if err != nil {
		log.Panicln(err)
	}
	err = p.Publish("My_NSQ_Topic", []byte("sample NSQ message"))
	if err != nil {
		log.Panicln(err)
	}
}
