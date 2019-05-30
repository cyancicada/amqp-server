package model

import (
	"github.com/go-xorm/xorm"
)

type (
	MpsModel struct {
		mpsMysql *xorm.Engine
	}
)

func NewMpsModel(mpsMysql *xorm.Engine) *MpsModel {

	return &MpsModel{mpsMysql: mpsMysql}
}

func (m *MpsModel) ExecSql(sqlorArgs []interface{}) error {
	_, err := m.mpsMysql.Transaction(func(session *xorm.Session) (i interface{}, e error) {
		return session.Exec(sqlorArgs...)
	})
	return err
}
