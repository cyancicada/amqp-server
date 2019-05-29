package handler

import (
	"database/sql"

	"ampp-server/common/rabbitmq"
	"ampp-server/common/utils"

	"github.com/go-xorm/xorm"
	"github.com/yakaa/log4g"
)

type (
	MpsHandler struct {
		mpsMysql *xorm.Engine
	}
)

func NewMpsHandler(mpsMysql *xorm.Engine) *MpsHandler {

	return &MpsHandler{mpsMysql: mpsMysql}
}

func (h *MpsHandler) Consumer(message *rabbitmq.Message) error {

	_, err := h.mpsMysql.Transaction(func(session *xorm.Session) (i interface{}, e error) {
		var err error
		var res sql.Result
		query := ""
		ks := ""
		placeholder := ""
		vs := []interface{}(nil)
		switch message.Operate {
		case rabbitmq.InsertType:
			ks, placeholder, vs = utils.SqlBuild(message.Data, ",")
			query = "insert into " + message.DataBase + "." + message.Table + " (" + ks + ") values (" + placeholder + ")"
		case rabbitmq.DeleteType:
			query = "delete from " + message.DataBase + "." + message.Table + " where " + message.Condition
		case rabbitmq.UpdateType:
			ks, _, vs = utils.SqlBuild(message.Data, "=?,")
			query = "update  " + message.DataBase + "." + message.Table + " sett " + ks + " where " + message.Condition
		}
		log4g.InfoFormat("Exec sql ===> %s ,ks is %+v, build data %+v", query, ks, vs)
		res, err = session.Exec(query, vs)
		return res, err
	})
	return err
}
