package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"ampp-server/common/rabbitmq"
	"ampp-server/common/rsa"
	"ampp-server/config"
	"ampp-server/model"
	"ampp-server/service"
	"github.com/streadway/amqp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
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
	mpsMysqlEngine, err := xorm.NewEngine("mysql", conf.MpsMysql.DataSource)
	erpMysqlEngine, err := xorm.NewEngine("mysql", conf.ErpMysql.DataSource)
	romeoMysqlEngine, err := xorm.NewEngine("mysql", conf.RomeoMysql.DataSource)

	mysqlEngine, err := xorm.NewEngine("mysql", conf.AmqpMysql.DataSource)
	if err != nil {
		log.Fatalf("read file %s: %s", *configFile, err)
	}
	amqpDial, err := amqp.Dial(conf.RabbitMq.DataSource)
	if err != nil {
		log.Fatalf("create connect fail %+v", err)
	}
	rsaObj, err := rsa.NewRsa(conf.RsaCert.PublicKeyPath, conf.RsaCert.PrivateKeyPath)
	if err != nil {
		log.Fatalf("create rsa fail %+v", err)
	}
	mpsConsumer := rabbitmq.BuildConsumer(
		amqpDial,
		conf.RabbitMq.MpsQueueName,
		service.NewMessageService(
			model.NewBaseModel(mpsMysqlEngine),
			model.NewMessagesModel(mysqlEngine)).ConsumerMessage,
	)
	erpConsumer := rabbitmq.BuildConsumer(
		amqpDial,
		conf.RabbitMq.ErpQueueName,
		service.NewMessageService(
			model.NewBaseModel(erpMysqlEngine),
			model.NewMessagesModel(mysqlEngine)).ConsumerMessage,
	)
	romeoConsumer := rabbitmq.BuildConsumer(
		amqpDial,
		conf.RabbitMq.RomeoQueueName,
		service.NewMessageService(
			model.NewBaseModel(romeoMysqlEngine),
			model.NewMessagesModel(mysqlEngine)).ConsumerMessage,
	)

	mpsConsumer.SetRsaRsaHelper(rsaObj)
	erpConsumer.SetRsaRsaHelper(rsaObj)
	romeoConsumer.SetRsaRsaHelper(rsaObj)

	defer rabbitmq.Close(amqpDial)
	rabbitmq.RunConsumes(erpConsumer, romeoConsumer, mpsConsumer)

}
