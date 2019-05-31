package utils

import (
	"strings"

	"ampp-server/common/rabbitmq"
)

func SqlBuild(data map[string]interface{}, sep string) (string, string, []interface{}) {
	ks := []string(nil)
	vs := []interface{}(nil)
	placeholder := "?,"
	for k, v := range data {
		ks = append(ks, k+sep)
		vs = append(vs, v)
	}
	return strings.Join(ks, ","),
		strings.Trim(strings.Repeat(placeholder, len(vs)), ","),
		vs
}

func ParseMessage(message *rabbitmq.Message) []interface{} {
	query := ""
	sqlorArgs := []interface{}(nil)
	switch message.Operate {
	case rabbitmq.InsertType:
		ks, placeholder, vs := SqlBuild(message.Data, "")
		query = "insert into " + message.DataBase + "." + message.Table + " (" + ks + ") values (" + placeholder + ")"
		sqlorArgs = append(sqlorArgs, query)
		sqlorArgs = append(sqlorArgs, vs...)
	case rabbitmq.DeleteType:
		query = "delete from " + message.DataBase + "." + message.Table + " where " + message.Condition
	case rabbitmq.UpdateType:
		ks, _, vs := SqlBuild(message.Data, "=?")
		query = "update  " + message.DataBase + "." + message.Table + " set " + ks + " where " + message.Condition
		sqlorArgs = append(sqlorArgs, query)
		sqlorArgs = append(sqlorArgs, vs...)
	}
	return sqlorArgs
}
