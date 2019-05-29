package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"ampp-server/common/rabbitmq"
	"ampp-server/config"
	"ampp-server/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yakaa/log4g"
)

var configFile = flag.String("f", "config/config.json", "Please set config file")

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

	mpsConsume, err := rabbitmq.NewConsumer(
		conf.MpsRabbitMq.DataSource,
		conf.MpsRabbitMq.QueueName,
	)
	if err != nil {
		log.Fatalf("build consume fail %+v", err)
	}
	defer mpsConsume.Close()
	mpsHandler := handler.NewMpsHandler(mpsMysqlEngine)
	log4g.Info("mps consumer start ...")
	log4g.Error(mpsConsume.StartConsume(mpsHandler.Consumer))
}
