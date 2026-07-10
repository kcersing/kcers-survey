package upload

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type FileUploadTask struct {
	FileName        string
	Cover           string
	FileTmpPath     string
	CoverTmpPath    string
	FileUploadPath  string
	CoverUploadPath string
}

type Publisher struct {
	exchange string
	ch       *amqp.Channel
}

// NewPublisher
func NewPublisher(conn *amqp.Connection, exchange string) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	if err := declareExchange(ch, exchange); err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}
	return &Publisher{
		ch:       ch,
		exchange: exchange,
	}, nil
}

func (p *Publisher) Publish(ctx context.Context, fileUploadTask FileUploadTask) error {
	body, err := json.Marshal(fileUploadTask)
	if err != nil {
		return fmt.Errorf("cannot marshal task: %v", err)
	}
	return p.ch.PublishWithContext(
		ctx,
		p.exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			//Headers:     amqp.Table{},
		})
}

// declareExchange
func declareExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil)
}

type Subscriber struct {
	conn     *amqp.Connection
	exchange string
}

// NewSubscriber
func NewSubscriber(conn *amqp.Connection, exchange string) (*Subscriber, error) {
	return &Subscriber{
		conn:     conn,
		exchange: exchange,
	}, nil
}

func getQueueName(queueName string) string {
	return queueName
}

// Subscribe
func (s *Subscriber) Subscribe(_ context.Context, queueName string, handler func([]byte) error) error {
	ch, err := s.conn.Channel()
	if err != nil {
		return fmt.Errorf("cannot allocate channel: %v", err.Error())
	}
	closeCh := func() {
		err := ch.Close()
		if err != nil {
			log.Printf("cannot close channel: %v", err.Error())
		}
	}

	if err := declareExchange(ch, s.exchange); err != nil {
		ch.Close()
		return fmt.Errorf("cannot declare exchange: %v", err)
	}
	fullQueueName := getQueueName(queueName)
	q, err := ch.QueueDeclare(
		fullQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("cannot declare queue: %s : %s", q.Name, err.Error())
	}
	if err := ch.QueueBind(q.Name, "", s.exchange, false, nil); err != nil {
		return fmt.Errorf("cannot bind queue: %s : %s", q.Name, err.Error())
	}
	cleanUp := func() {
		_, err := ch.QueueDelete(q.Name, false, false, false)
		if err != nil {
			log.Printf("cannot delete queue: %s : %s", q.Name, err.Error())
		}
		closeCh()
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("cannot consume queue: %s : %s", q.Name, err.Error())
	}

	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				cleanUp()
				return fmt.Errorf("channel closed: %s ", q.Name)
			}
			if len(msg.Body) == 0 {
				if msg.Acknowledger != nil {
					if err = msg.Nack(false, false); err != nil {
						log.Printf("nack empty msg err: %v", err)
					}
				}
				continue
			}

			if err = handler(msg.Body); err != nil {
				log.Printf("handler err: %v", err)
				if msg.Acknowledger != nil {
					if nackErr := msg.Nack(false, true); nackErr != nil {
						log.Printf("nack err: %v", nackErr)
					}
				}
				continue
			}
			if msg.Acknowledger != nil {
				if ackErr := msg.Ack(false); ackErr != nil {
					return fmt.Errorf("ack fail: %s", q.Name)
				}
			}
		}

	}

}
