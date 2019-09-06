package rabbitmq

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

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

var deadSignal = []os.Signal{
	syscall.SIGTERM,
	syscall.SIGINT,
	syscall.SIGKILL,
	syscall.SIGHUP,
	syscall.SIGQUIT,
}

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

func (c *Consumer) Run() error {
	c.Close()
	if err := c.StartConsume(); err != nil {
		return err
	}
	return nil
}

func (c *Consumer) Close() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, deadSignal...)
	go func() {
		log4g.InfoFormat(" receive dead signal %+v ", <-ch)
		if err := c.amqpDial.Close(); err != nil {
			log4g.InfoFormat("Consumer conn Close err %+v by receive dead signal", err)
		}
		c.stop <- true
		os.Exit(1)
	}()
}
