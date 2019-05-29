package handler

import (
	"ampp-server/common/rabbitmq"

	"github.com/go-xorm/xorm"
)

type (
	ErpHandler struct {
		erpMysql *xorm.Engine
	}
)

func NewErpHandler(erpMysql *xorm.Engine) (*ErpHandler) {

	return &ErpHandler{erpMysql: erpMysql}
}

func (h *ErpHandler) Consumer(message *rabbitmq.Message) error {

	return nil
}
