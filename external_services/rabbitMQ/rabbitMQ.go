package rabbitmq

import (
	"architecture_template/constants/notis"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type rabbitMQClient struct {
	logger  *log.Logger
	channel *amqp.Channel
	cnn     *amqp.Connection
}

var rbClient *rabbitMQClient

func GetRabbitMQClient(cnnStr string, logger *log.Logger) (*rabbitMQClient, error) {
	if rbClient != nil {
		return rbClient, nil
	}

	res, err := initializeRabbitMQClient(cnnStr, logger)
	if err != nil {
		return nil, err
	}

	rbClient = res
	return res, nil
}

func initializeRabbitMQClient(cnnStr string, logger *log.Logger) (*rabbitMQClient, error) {
	cnn, err := amqp.Dial(cnnStr)
	var internalErr error = errors.New(notis.InternalErr)
	var prefixErrMsg string = fmt.Sprintf(notis.RabbitmqConnectMsg, cnnStr)

	if err != nil {
		logger.Print(prefixErrMsg + err.Error())
		return nil, internalErr
	}

	channel, err := cnn.Channel()
	if err != nil {
		logger.Print(prefixErrMsg + err.Error())
		return nil, internalErr
	}

	return &rabbitMQClient{
		logger:  logger,
		channel: channel,
		cnn:     cnn,
	}, nil
}

func (client *rabbitMQClient) GetChannel() *amqp.Channel {
	return client.channel
}

func (client *rabbitMQClient) Publish(queue string, data interface{}) error {
	jsonData, _ := json.Marshal(data)
	//var internalErr error = errors.New(notis.InternalErr)

	if err := client.channel.Publish(
		queue,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		},
	); err != nil {
		client.logger.Print(fmt.Sprintf(notis.RabbitmqPublishMsg, queue) + err.Error())
		return errors.New(notis.InternalErr)
	}

	return nil
}

func (client *rabbitMQClient) Consume(queue string) (<-chan amqp.Delivery, error) {
	res, err := client.channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		client.logger.Print(fmt.Sprintf(notis.RabbitmqConsumeMsg, queue) + err.Error())
		return nil, errors.New(notis.InternalErr)
	}

	return res, nil
}

func (client *rabbitMQClient) Declare(queue string) error {
	_, err := client.channel.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		client.logger.Print(fmt.Sprintf(notis.RabbitmqDeclareMsg, queue) + err.Error())
		return errors.New(notis.InternalErr)
	}

	return nil
}
