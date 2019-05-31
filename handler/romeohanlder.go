package handler

import (
	"ampp-server/common/rabbitmq"
	"ampp-server/service"
)

type (
	RomeoHandler struct {
		romeoService *service.RomeoService
	}
)

func NewRomeoHandler(romeoService *service.RomeoService) *RomeoHandler {

	return &RomeoHandler{romeoService: romeoService}
}

func (h *RomeoHandler) Consumer(message *rabbitmq.Message) error {
	return h.romeoService.OperateMessage(message)
}
