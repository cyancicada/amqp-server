package service

import (
	"encoding/json"

	"ampp-server/common/httpx"
	"ampp-server/common/rabbitmq"
	"ampp-server/common/utils"
	"ampp-server/model"

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

	status := model.SuccessMessageStatus
	var err error
	var responseStatus bool
	switch message.Operate {
	case rabbitmq.HttpType:
		if responseStatus, err = utils.HttpRequest(httpx.HttpMethodPost, message.Url, message.Data); err != nil {
			log4g.ErrorFormat("utils.HttpRequest  %+v  %+v", message, err)
		}
	default:
		return nil
	}
	if err == nil || responseStatus {
		return nil
	}
	if bs, err := json.Marshal(message); err == nil {
		status = model.FailMessageStatus
		if _, err := s.messageModel.Insert(&model.Messages{
			Message: string(bs),
			Status:  status,
		}); err != nil {
			log4g.ErrorFormat("ConsumerMessage s.messageModel.Insert Err %+v", err)
		}
	}
	return nil
}
