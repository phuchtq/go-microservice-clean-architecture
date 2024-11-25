package helper

import (
	"architecture_template/constants/notis"
	rabbitmq "architecture_template/external_services/rabbitMQ"
	"errors"
	"log"

	"github.com/streadway/amqp"
)

func InitializeRabbitMQQueue(cnnStr string, logger *log.Logger, queues []string) error {
	client, err := rabbitmq.GetRabbitMQClient(cnnStr, logger)

	if err != nil {
		return err
	}

	for _, queue := range queues {
		if err := client.Declare(queue); err != nil {
			logger.Print()
			return errors.New(notis.InternalErr)
		}
	}

	return nil
}

func UtilizeMessage[T any](cnnStr, queue string, logger *log.Logger, method func(data T) error) error {
	client, err := rabbitmq.GetRabbitMQClient(cnnStr, logger)
	if err != nil {
		return err
	}

	var internalErr error = errors.New(notis.InternalErr)

	if err := client.Declare(queue); err != nil {
		logger.Print()
		return internalErr
	}

	msgs, err := client.Consume(queue)
	if err != nil {
		logger.Print()
		return internalErr
	}

	extractData, err := extractMessage[T](msgs)
	return method(*extractData)

}

func extractMessage[T any](msgs <-chan amqp.Delivery) (*T, error) {
	for delivery := range msgs {
		var res = ConvertJsonToModel[T](string(delivery.Body))
		if res != nil {
			return res, nil
		}
	}

	return nil, nil
}
