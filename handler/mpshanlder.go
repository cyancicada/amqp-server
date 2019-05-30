package handler

import (
	"ampp-server/common/rabbitmq"
	"ampp-server/service"
)

type (
	MpsHandler struct {
		mpsService *service.MpsService
	}
)

func NewMpsHandler(mpsService *service.MpsService) *MpsHandler {

	return &MpsHandler{mpsService: mpsService}
}

func (h *MpsHandler) Consumer(message *rabbitmq.Message) error {
	return h.mpsService.OperateMessage(message)
}
