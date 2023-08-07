package broker

import (
	"errors"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var rabbitMQConn *amqp.Connection

func InitRabbitMQConnection(connString string) {
	var err error
	retries := 5
	for i := 0; i < retries; i++ {
		rabbitMQConn, err = amqp.Dial(connString)
		if err == nil {
			return
		}
		log.Printf("Failed to connect to RabbitMQ, attempt %d/%d: %v", i+1, retries, err)
		time.Sleep(5 * time.Second)
	}
	log.Fatalf("Failed to connect to RabbitMQ after %d attempts: %v", retries, err)
}

func PublishMessageWithCorrelationID(queueName, replyQueue, message, correlationID string) error {
	if rabbitMQConn == nil {
		return errors.New("RabbitMQ connection is not initialized")
	}
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		return err
	}
	//defer func() {
	//	if cerr := ch.Close(); cerr != nil {
	//		log.Printf("Error closing the channel: %v", cerr)
	//	}
	//}()

	msg := amqp.Publishing{
		ContentType:   "application/json",
		Body:          []byte(message),
		ReplyTo:       replyQueue,
		CorrelationId: correlationID,
	}

	return ch.Publish("", queueName, false, false, msg)
}

func ConsumeMessageWithCorrelationID(queueName string) (string, string, error) {
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		return "", "", err
	}
	//defer func() {
	//	if cerr := ch.Close(); cerr != nil {
	//		log.Printf("Error closing the channel: %v", cerr)
	//	}
	//}()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return "", "", err
	}

	select {
	case msg := <-msgs:
		return string(msg.Body), msg.CorrelationId, nil
	case <-time.After(time.Second * 10): // timeout after 10 seconds
		return "", "", errors.New("Timeout waiting for response")
	}
}

func DeclareQueue(queueName string) error {
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		return err
	}
	//defer func() {
	//	if cerr := ch.Close(); cerr != nil {
	//		log.Printf("Error closing the channel: %v", cerr)
	//	}
	//}()
	_, err = ch.QueueDeclare(
		queueName, // name of the queue
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return err
}
