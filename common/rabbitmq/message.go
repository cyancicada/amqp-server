package rabbitmq

import "time"

type (
	Message struct {
		Data      map[string]interface{} `json:"data"`
		Timestamp time.Time              `json:"timestamp"`
		Type      string                 `json:"type"`
		Url       string                 `json:"url"`
		Retry     bool                   `json:"retry"`
	}
)

const (
	ContentType string = "Content-Type:application/json"
	HttpType    string = "HTTP"
)
