package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"github.com/yakaa/log4g"
)

type (
	Consumer struct {
		amqpDial     *amqp.Connection
		queueName    string
		ConsumerName string
	}
	ConsumerFunc func(message *Message) error
)

func NewConsumer(dataSource, queueName string) (*Consumer, error) {
	amqpDial, err := amqp.Dial(dataSource)
	if err != nil {
		return nil, err
	}
	return &Consumer{amqpDial: amqpDial, queueName: queueName}, nil
}

func (p *Consumer) SetConsumerName(consumerName string) {
	p.ConsumerName = consumerName
}

func (p *Consumer) StartConsume(consumerFunc ConsumerFunc) error {
	forever := make(chan bool)
	ch, err := p.amqpDial.Channel()
	if err != nil {
		return err
	}
	//if err = ch.Qos(1, 0, false); err != nil {
	//	return err
	//}
	defer func() {
		log4g.ErrorFormat("Publish Close Ch err %+v", ch.Close())
	}()
	response, err := ch.Consume(p.queueName, p.ConsumerName, false, false, false, false, nil)
	go func() {
		for d := range response {
			message := new(Message)
			if err := json.Unmarshal(d.Body, message); err != nil {
				log4g.ErrorFormat("Err Message format %+v", err)
				if err := d.Ack(false); err != nil {
					log4g.ErrorFormat("d.Ack message fail err %+v", err)
				}
			} else {
				if err := consumerFunc(message); err != nil {
					log4g.ErrorFormat("Consume Message Err %+v", err)
					log4g.InfoFormat("ch.Reject Error %+v", ch.Reject(d.DeliveryTag, true))
				} else {
					if err := d.Ack(false); err != nil {
						log4g.ErrorFormat("d.Ack message fail err %+v", err)
					}
				}
			}
		}
	}()
	<-forever
	return nil
}

func (p *Consumer) Close() {
	if err := p.amqpDial.Close(); err != nil {
		log4g.ErrorFormat("Consumer conn Close err %+v", err)
	}
}
