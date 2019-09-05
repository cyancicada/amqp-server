package model

type (
	MessagesModel struct {
	}
)

const (
	SuccessMessageStatus string = "SUCCESS"
	FailMessageStatus    string = "FAIL"
)

func NewMessagesModel() *MessagesModel {

	return new(MessagesModel)
}
