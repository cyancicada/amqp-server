package rabbitmq

import "time"

type (
	OperateType string
	Message     struct {
		Operate   OperateType            `json:"operate"`
		Data      map[string]interface{} `json:"data"`
		Condition string                 `json:"condition"`
		Timestamp time.Time              `json:"timestamp"`
		Type      string                 `json:"type"`
		Url       string                 `json:"url"`
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
