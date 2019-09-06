package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"yasuo/common/rabbitmq"
	"yasuo/config"
	"yasuo/model"
	"yasuo/service"

	"github.com/streadway/amqp"
	"github.com/yakaa/log4g"
)

var configFile = flag.String("c", "config/config.json", "Please set config file")

func main() {
	flag.Parse()
	body, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("read file %s: %s", *configFile, err)
	}
	conf := new(config.Config)
	if err := json.Unmarshal(body, conf); err != nil {
		log.Fatalf("json.Unmarshal %s: %s", *configFile, err)
	}
	log4g.Init(conf.Log4g)

	amqpDial, err := amqp.Dial(conf.RabbitMq.DataSource)
	if err != nil {
		log.Fatalf("create connect fail %+v", err)
	}

	consumer := rabbitmq.BuildConsumer(
		amqpDial,
		conf.RabbitMq.QueueName,
		service.NewMessageService(model.NewMessagesModel()).ConsumerMessage,
	)
	log4g.Error(consumer.Run())

}
