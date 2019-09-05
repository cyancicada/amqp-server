package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"github.com/yakaa/log4g"
	"yasuo/common/rsa"
)

type (
	Consumer struct {
		amqpDial     *amqp.Connection
		queueName    string
		ConsumerName string
		stop         chan bool
		consumerFunc ConsumerFunc
		rsaHelper    *rsa.Rsa
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
	response, err := ch.Consume(c.queueName, c.ConsumerName, true, false, false, false, nil)
	go func() {
		for d := range response {
			message := new(Message)
			if err := json.Unmarshal(d.Body, message); err != nil {
				log4g.ErrorFormat("Err Message format %+v", err)
				continue
			}
			if err := c.consumerFunc(message); err != nil {
				log4g.ErrorFormat("Consume Message Err %+v", err)
				continue
			}
			log4g.InfoFormat("Consume message Success  %+v", message)
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

func (c *Consumer) GetRsaHelper() *rsa.Rsa {
	return c.rsaHelper
}

func (c *Consumer) SetRsaRsaHelper(rsaHelper *rsa.Rsa) {
	c.rsaHelper = rsaHelper
}

func Close(amqpDial *amqp.Connection) {
	if err := amqpDial.Close(); err != nil {
		log4g.ErrorFormat("Consumer conn Close err %+v", err)
	}
}
