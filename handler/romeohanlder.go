package handler

import (
	"ampp-server/common/rabbitmq"

	"github.com/go-xorm/xorm"
)

type (
	RomeoHandler struct {
		romeoMysql *xorm.Engine
	}
)

func NewRomeoHandler(mpsMysql *xorm.Engine) (*RomeoHandler) {

	return &RomeoHandler{romeoMysql: mpsMysql}
}

func (h *RomeoHandler) Consumer(message *rabbitmq.Message) error {

	return nil
}
