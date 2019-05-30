package model

import "github.com/go-xorm/xorm"

type (
	MessagesModel struct {
		mpsMysql *xorm.Engine
	}

	Messages struct {
		Id      int64
		Message string `xorm:"varchar(255) 'message'" json:"message"`
		Status  string `xorm:"varchar(10) 'status'" json:"status"`
	}
)

const (
	SuccessMessageStatus string = "SUCCESS"
	FailMessageStatus    string = "FAIL"
)

func NewMessagesModel(mpsMysql *xorm.Engine) *MessagesModel {

	return &MessagesModel{mpsMysql: mpsMysql}
}

func (m *MessagesModel) Insert(message *Messages) (int64, error) {
	return m.mpsMysql.InsertOne(message)
}
