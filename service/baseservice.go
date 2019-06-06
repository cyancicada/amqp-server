package service

import "ampp-server/common/rabbitmq"

type (
	BaseMessageConsumerService interface {
		ConsumerMessage(message *rabbitmq.Message) error
	}
)
