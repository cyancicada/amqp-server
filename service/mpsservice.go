package service

import (
	"encoding/json"

	"ampp-server/common/rabbitmq"
	"ampp-server/common/utils"
	"ampp-server/model"
	"github.com/yakaa/log4g"
)

type (
	MpsService struct {
		mpsModel     *model.MpsModel
		messageModel *model.MessagesModel
	}
)

func NewMpsService(mpsModel *model.MpsModel, messageModel *model.MessagesModel) *MpsService {

	return &MpsService{mpsModel: mpsModel, messageModel: messageModel}
}

func (s *MpsService) OperateMessage(message *rabbitmq.Message) error {
	query := ""
	sqlorArgs := []interface{}(nil)
	switch message.Operate {
	case rabbitmq.InsertType:
		ks, placeholder, vs := utils.SqlBuild(message.Data, "")
		query = "insert into " + message.DataBase + "." + message.Table + " (" + ks + ") values (" + placeholder + ")"
		sqlorArgs = append(sqlorArgs, query)
		sqlorArgs = append(sqlorArgs, vs...)
	case rabbitmq.DeleteType:
		query = "delete from " + message.DataBase + "." + message.Table + " where " + message.Condition
	case rabbitmq.UpdateType:
		ks, _, vs := utils.SqlBuild(message.Data, "=?")
		query = "update  " + message.DataBase + "." + message.Table + " set " + ks + " where " + message.Condition
		sqlorArgs = append(sqlorArgs, query)
		sqlorArgs = append(sqlorArgs, vs...)
	}
	status := model.SuccessMessageStatus
	if err := s.mpsModel.ExecSql(sqlorArgs); err != nil {
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
