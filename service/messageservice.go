package service

import (
	"yasuo/common/httpx"
	"yasuo/common/rabbitmq"
	"yasuo/common/utils"
	"yasuo/model"

	"github.com/yakaa/log4g"
)

type (
	MessageService struct {
		messageModel *model.MessagesModel
	}
)

func NewMessageService(messageModel *model.MessagesModel) *MessageService {

	return &MessageService{messageModel: messageModel}
}

func (s *MessageService) ConsumerMessage(message *rabbitmq.Message) error {
	var err error
	var responseStatus bool
	if responseStatus, err = utils.HttpRequest(httpx.HttpMethodPost, message.Url, message.Data); err != nil {
		log4g.ErrorFormat("utils.HttpRequest  %+v  %+v", message, err)
	}
	if err == nil || responseStatus {
		return nil
	}
	return err
}
