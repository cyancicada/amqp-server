package service

import "yasuo/common/rabbitmq"

type (
	BaseMessageConsumerService interface {
		ConsumerMessage(message *rabbitmq.Message) error
	}
)
