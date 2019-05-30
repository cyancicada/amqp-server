package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"github.com/yakaa/log4g"
)

type (
	Publisher struct {
		amqpDial  *amqp.Connection
		queueName string
		Exchange  string
	}
)

func NewPublisher(dataSource, queueName string) (*Publisher, error) {
	amqpDial, err := amqp.Dial(dataSource)
	if err != nil {
		return nil, err
	}
	return &Publisher{amqpDial: amqpDial, queueName: queueName}, nil
}

func (p *Publisher) SetExchange(exchange string) {
	p.Exchange = exchange
}

func (p *Publisher) Push(message Message) error {
	ch, err := p.amqpDial.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			log4g.ErrorFormat("Publish Close Ch err %+v", err)
		}
	}()
	q, err := ch.QueueDeclare(p.queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	if err = ch.Publish(p.Exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         body,
	}); err != nil {
		return err
	}
	return nil
}

func (p *Publisher) Close() {
	if err := p.amqpDial.Close(); err != nil {
		log4g.ErrorFormat("Publisher conn Close err %+v", err)
	}
}
