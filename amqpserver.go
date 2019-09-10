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
	consumer, err := rabbitmq.BuildConsumer(
		conf.RabbitMq,
		service.NewMessageService(model.NewMessagesModel()).ConsumerMessage,
	)
	if err != nil {
		log.Fatalf("create publisher fail %+v", err)
	}
	log4g.Error(consumer.Run())
}
