package handler

import (
	"ampp-server/common/rabbitmq"
	"ampp-server/service"
)

type (
	ErpHandler struct {
		erpService *service.ErpService
	}
)

func NewErpHandler(erpService *service.ErpService) *ErpHandler {

	return &ErpHandler{erpService: erpService}
}

func (h *ErpHandler) Consumer(message *rabbitmq.Message) error {
	return h.erpService.OperateMessage(message)
}
