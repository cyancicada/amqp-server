package rabbitmq

import "time"

type (
	OperateType string
	Message     struct {
		DataBase   string                 `json:"dataBase"`
		Table      string                 `json:"table"`
		Operate    OperateType            `json:"operate"`
		Data       map[string]interface{} `json:"data"`
		Condition  string                 `json:"condition"`
		Expiration string                 `json:"expiration"`
		MessageId  string                 `json:"messageId"`
		Timestamp  time.Time              `json:"timestamp"`
		Type       string                 `json:"type"`
		UserId     string                 `json:"userId"`
		AppId      string                 `json:"appId"`
		Url        string                 `json:"url"`
	}
)

const (
	ContentType string      = "Content-Type:application/json"
	HttpType    OperateType = "HTTP"
	InsertType  OperateType = "INSERT"
	UpdateType  OperateType = "UPDATE"
	DeleteType  OperateType = "DELETE"
	SelectType  OperateType = "SELECT"
)
