package broker

import (
	"encoding/json"
	"errors"
	"gateway_processor/protos"
	"github.com/google/uuid"
	"log"
	"time"
)

func GenerateUniqueID() string {
	return uuid.New().String()
}

func PublishToQueue(req *protos.CreateUserRequest, publishQueue, responseQueue, correlationID string) error {
	messageJSON, err := json.Marshal(req)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return err
	}

	err = PublishMessageWithCorrelationID(
		publishQueue, responseQueue, string(messageJSON), correlationID,
	)
	if err != nil {
		log.Println("Error publishing message:", err)
		return err
	}

	return nil
}

func WaitForResponse(responseQueue, correlationID string) (*protos.User, error) {
	responseCh := make(chan *protos.User)
	errorCh := make(chan error)

	// Ждем ответа в горутине, чтобы не блокировать исполнение
	go func() {
		response, receivedCorrelationID, err := ConsumeMessageWithCorrelationID(responseQueue)
		if err != nil {
			errorCh <- err
			return
		}

		if receivedCorrelationID != correlationID {
			errorCh <- errors.New("Mismatched correlation ID")
			return
		}

		var user protos.User
		err = json.Unmarshal([]byte(response), &user)
		if err != nil {
			errorCh <- err
			return
		}

		responseCh <- &user
	}()

	select {
	case res := <-responseCh:
		return res, nil
	case err := <-errorCh:
		return nil, err
	case <-time.After(time.Second * 10): // timeout after 10 seconds
		return nil, errors.New("Timeout waiting for response")
	}
}
