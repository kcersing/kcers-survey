package upload

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
}

func Init(cfg Config) *Service {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	amqpConn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("cannot dial amqp: %v", err.Error())
	}
	publisher, err := NewPublisher(amqpConn, "exchange")
	if err != nil {
		log.Printf("cannot create publisher: %v", err.Error())
	}
	subscriber, err := NewSubscriber(amqpConn, "exchange")
	if err != nil {
		log.Printf("cannot create subscriber: %v", err.Error())
	}

	uploader := &FastdfsUploader{}

	return NewUploadService(publisher, subscriber, uploader)
}
