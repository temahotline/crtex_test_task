package broker

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

var conn *amqp.Connection
var ch *amqp.Channel

func InitRabbitMQConnection(connString string) {
	var err error
	retries := 5
	for i := 0; i < retries; i++ {
		conn, err = amqp.Dial(connString)
		if err == nil {
			return
		}
		log.Printf("Failed to connect to RabbitMQ, attempt %d/%d: %v", i+1, retries, err)
		time.Sleep(5 * time.Second)
	}
	log.Fatalf("Failed to connect to RabbitMQ after %d attempts: %v", retries, err)
}

func ConsumeMessages(queueName string) {
	log.Println("Consuming messages!")
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}
	log.Println(msgs)

	//go func() {
	//	for d := range msgs {
	//		log.Printf("Received a message: %s", d.Body)
	//		handleIncomingMessage(string(d.Body), d.ReplyTo, d.CorrelationId)
	//	}
	//}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever
}

func PublishMessage(queueName, message, correlationID string) error {
	msg := amqp.Publishing{
		ContentType:   "application/json",
		Body:          []byte(message),
		CorrelationId: correlationID,
	}

	return ch.Publish("", queueName, false, false, msg)
}

func handleIncomingMessage(message, replyTo, correlationID string) {
	log.Println("Handling incoming message!")
	// Здесь ваш код для обработки сообщения.
	// Например, парсинг JSON, сохранение данных в базу данных и формирование ответа.

	response := "Ваш ответ на сообщение" // Вы можете сформировать этот ответ на основе данных из базы данных

	err := PublishMessage(replyTo, response, correlationID)
	if err != nil {
		log.Printf("Error sending response: %v", err)
	}
}
