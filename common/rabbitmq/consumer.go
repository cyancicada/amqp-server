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
		stop         chan bool
		consumerFunc ConsumerFunc
	}
	ConsumerFunc func(message *Message) error
)

func BuildConsumer(amqpDial *amqp.Connection, queueName string, consumerFunc ConsumerFunc) *Consumer {
	return &Consumer{amqpDial: amqpDial, queueName: queueName, stop: make(chan bool), consumerFunc: consumerFunc}
}

func (c *Consumer) SetConsumerName(consumerName string) {
	c.ConsumerName = consumerName
}

func (c *Consumer) StartConsume() error {
	ch, err := c.amqpDial.Channel()
	if err != nil {
		return err
	}
	if err = ch.Qos(1, 0, false); err != nil {
		return err
	}
	defer func() {
		log4g.ErrorFormat("Consumer Close Ch err %+v", ch.Close())
	}()
	response, err := ch.Consume(c.queueName, c.ConsumerName, false, false, false, false, nil)
	go func() {
		for d := range response {
			message := new(Message)
			if err := json.Unmarshal(d.Body, message); err != nil {
				log4g.ErrorFormat("Err Message format %+v", err)
				if err := d.Ack(false); err != nil {
					log4g.ErrorFormat("d.Ack message fail err %+v", err)
				}
			} else {
				if err := c.consumerFunc(message); err != nil {
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
	<-c.stop
	return nil
}

func RunConsumes(consumers ...*Consumer) {
	if len(consumers) == 0 {
		return
	}
	forever := make(chan bool)
	for _, consumer := range consumers {
		go func(c *Consumer) {
			log4g.InfoFormat("start Consumer queueName [%s]...", c.queueName)
			if err := c.StartConsume(); err != nil {
				log4g.ErrorFormat("consumer.StartConsume fail %+v", err)
			}
		}(consumer)
	}
	<-forever
}

func Close(amqpDial *amqp.Connection) {
	if err := amqpDial.Close(); err != nil {
		log4g.ErrorFormat("Consumer conn Close err %+v", err)
	}
}
