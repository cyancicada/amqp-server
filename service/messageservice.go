package service

import (
	"encoding/json"
	"strings"

	"ampp-server/common/rabbitmq"
	"ampp-server/common/utils"
	"ampp-server/model"

	"github.com/yakaa/log4g"
)

type (
	MessageService struct {
		baseModel    *model.BaseModel
		messageModel *model.MessagesModel
	}
)

func NewMessageService(baseModel *model.BaseModel, messageModel *model.MessagesModel) *MessageService {

	return &MessageService{baseModel: baseModel, messageModel: messageModel}
}

func (s *MessageService) ConsumerMessage(message *rabbitmq.Message) error {
	log4g.ErrorFormat("utils.Execute message ====> %+v", message)
	status := model.SuccessMessageStatus
	var err error
	var param []byte
	var requestStatus string
	switch message.Operate {
	case rabbitmq.CurlType:
		param, err = json.Marshal(message.Data)
		if err != nil {
			return nil
		}
		curlArgs := append(rabbitmq.CurlRunParamArray, string(param), message.Url)
		if requestStatus, err = utils.Execute(strings.ToLower(string(rabbitmq.CurlType)), curlArgs...); err != nil {
			log4g.ErrorFormat("utils.Execute curlArgs %+v  %+v", curlArgs, err)
		}
	default:
		return nil
	}
	if err == nil || requestStatus == rabbitmq.MessageConsumeSuccessStatus {
		return nil
	}
	if bs, err := json.Marshal(message); err == nil {
		if _, err := s.messageModel.Insert(&model.Messages{
			Message: string(bs),
			Status:  status,
		}); err != nil {
			log4g.ErrorFormat("ConsumerMessage s.messageModel.Insert Err %+v", err)
		}
	}
	return nil
}
