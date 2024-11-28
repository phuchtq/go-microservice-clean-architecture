package helper

import (
	rabbitmq "architecture_template/external_services/rabbitMQ"
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
			return err
		}
	}

	return nil
}

func UtilizeMessage[T any](cnnStr, queue string, logger *log.Logger, data T, method func(extractMessage T) error) error {
	client, err := rabbitmq.GetRabbitMQClient(cnnStr, logger)
	if err != nil {
		return err
	}

	if err := client.Declare(queue); err != nil {
		return err
	}

	msgs, err := client.Consume(queue)
	if err != nil {
		return err
	}

	extractData, err := extractMessage(msgs, data)

	if err != nil {
		return err
	}

	return method(*extractData)
}

func PublishEvent(cnnStr string, queue string, logger *log.Logger, data interface{}) error {
	client, err := rabbitmq.GetRabbitMQClient(cnnStr, logger)
	if err != nil {
		return err
	}

	if err := client.Declare(queue); err != nil {
		return err
	}

	return client.Publish(queue, data)
}

func extractMessage[T any](msgs <-chan amqp.Delivery, data T) (*T, error) {
	for delivery := range msgs {
		var res = ConvertJsonToModel[T](string(delivery.Body))
		if res != nil && res == &data {
			return res, nil
		}
	}

	return nil, nil
}
