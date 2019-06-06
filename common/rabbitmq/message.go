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

//"curl",
//		"-H",
//		"Content-Type:application/json",
//		"-X",
//		"POST",
//		"-d",
//		`{"user": "admin", "passwd":"12345678"}`,
//		"http://www.xxx.com/test.php",

var (
	CurlRunParamArray = []string{"-H", "Content-Type:application/json", "-X", "POST", "-d"}
)

const (
	ContentType string      = "Content-Type:application/json"
	CurlType    OperateType = "CURL"
	InsertType  OperateType = "INSERT"
	UpdateType  OperateType = "UPDATE"
	DeleteType  OperateType = "DELETE"
	SelectType  OperateType = "SELECT"

	MessageConsumeSuccessStatus string = "SUCCESS"
	MessageConsumeFailStatus    string = "FAIL"
)
