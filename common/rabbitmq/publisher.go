package rabbitmq

import (
	"encoding/json"
	"os"
	"os/signal"

	"github.com/streadway/amqp"
	"github.com/yakaa/log4g"
	"yasuo/config"
)

type (
	Publisher struct {
		amqpDial   *amqp.Connection
		amqpDialCh *amqp.Channel
		conf       config.RabbitMq
	}
)

func BuildPublisher(conf config.RabbitMq) (*Publisher, error) {
	amqpDial, err := amqp.Dial(conf.DataSource)
	if err != nil {
		return nil, err
	}
	ch, err := amqpDial.Channel()
	if err != nil {
		return nil, err
	}
	return &Publisher{amqpDial: amqpDial, conf: conf, amqpDialCh: ch}, nil
}

func (p *Publisher) Push(message *Message) error {
	q, err := p.amqpDialCh.QueueDeclare(
		p.conf.QueueName,
		p.conf.Durable,
		p.conf.AutoDelete,
		p.conf.Exclusive,
		p.conf.NoWait,
		p.conf.Args,
	)
	if err != nil {
		return err
	}
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	if err = p.amqpDialCh.Publish(p.conf.Exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         body,
	}); err != nil {
		return err
	}
	return nil
}

func (p *Publisher) GetQueueName() string {
	return p.conf.QueueName
}

func (p *Publisher) Close() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, DeadSignal...)
	go func() {
		log4g.InfoFormat("Publisher receive dead signal %+v ", <-ch)
		if err := p.amqpDialCh.Close(); err != nil {
			log4g.ErrorFormat("c.amqpDialCh.Close err %+v", err)
		}
		if err := p.amqpDial.Close(); err != nil {
			log4g.InfoFormat("Publisher conn Close err %+v by receive dead signal", err)
		}
		os.Exit(1)
	}()
}
