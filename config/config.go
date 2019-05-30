package config

import "github.com/yakaa/log4g"

//amqp.Dial("amqp://guest:guest@localhost:5672/")
type (
	Config struct {
		Log4g         log4g.Config
		ErpRabbitMq   RabbitMq
		MpsRabbitMq   RabbitMq
		RomeoRabbitMq RabbitMq
		MpsMysql      Mysql
		AmqpMysql     Mysql
		ErpMysql      Mysql
		RomeoMysql    Mysql
	}

	RabbitMq struct {
		DataSource string
		QueueName  string
	}
	Mysql struct {
		DataSource string
	}
)
