package service

import (
	"encoding/json"

	"ampp-server/common/rabbitmq"
	"ampp-server/common/utils"
	"ampp-server/model"

	"github.com/yakaa/log4g"
)

type (
	ErpService struct {
		baseModel    *model.BaseModel
		messageModel *model.MessagesModel
	}
)

func NewErpService(baseModel *model.BaseModel, messageModel *model.MessagesModel) *ErpService {

	return &ErpService{baseModel: baseModel, messageModel: messageModel}
}

func (s *ErpService) OperateMessage(message *rabbitmq.Message) error {
	status := model.SuccessMessageStatus
	if err := s.baseModel.ExecSql(utils.ParseMessage(message)); err != nil {
		log4g.ErrorFormat("OperateMessage s.mpsModel.ExecSql Err %+v", err)
		status = model.FailMessageStatus
	}
	if bs, err := json.Marshal(message); err == nil {
		if _, err := s.messageModel.Insert(&model.Messages{
			Message: string(bs),
			Status:  status,
		}); err != nil {
			log4g.ErrorFormat("OperateMessage s.messageModel.Insert Err %+v", err)
		}
	}
	return nil
}
