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

//
//func main() {
//	flag.Parse()
//	body, err := ioutil.ReadFile(*configFile)
//	if err != nil {
//		log.Fatalf("read file %s: %s", *configFile, err)
//	}
//	conf := new(config.Config)
//	if err := json.Unmarshal(body, conf); err != nil {
//		log.Fatalf("json.Unmarshal %s: %s", *configFile, err)
//	}
//	mpsPublisher, err := rabbitmq.NewPublisher(
//		conf.MpsRabbitMq.DataSource,
//		conf.MpsRabbitMq.QueueName,
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	i := 17
//	for {
//		operate := rabbitmq.InsertType
//		switch i % 3 {
//		case 0:
//			operate = rabbitmq.UpdateType
//		case 1:
//			operate = rabbitmq.DeleteType
//
//		}
//		log4g.Info(mpsPublisher.Push(rabbitmq.Message{
//			DataBase: "mps",
//			Table:    "AuthItem",
//			Operate:  operate,
//			Data: map[string]interface{}{
//				"name":        "yk_" + strconv.Itoa(i),
//				"type":        "type_" + strconv.Itoa(i),
//				"description": "description" + strconv.Itoa(i),
//				"bizrule":     "bizrule_" + strconv.Itoa(i),
//				"data":        "data_" + strconv.Itoa(i),
//			},
//		}));
//		time.Sleep(3 * time.Second)
//	}
//
//	//CREATE TABLE `AuthItem` (
//	//  `name` varchar(64) NOT NULL,
//	//  `type` int(11) NOT NULL,
//	//  `description` text,
//	//  `bizrule` text,
//	//  `data` text,
//	//  PRIMARY KEY (`name`) USING BTREE
//	//) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
//}
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
