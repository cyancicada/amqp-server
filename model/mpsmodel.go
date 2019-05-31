package model

import (
	"github.com/go-xorm/xorm"
)

type (
	BaseModel struct {
		mysqlEngine *xorm.Engine
	}
)

func NewBaseModel(mysqlEngine *xorm.Engine) *BaseModel {

	return &BaseModel{mysqlEngine: mysqlEngine}
}

func (m *BaseModel) ExecSql(sqlorArgs []interface{}) error {
	_, err := m.mysqlEngine.Transaction(func(session *xorm.Session) (i interface{}, e error) {
		return session.Exec(sqlorArgs...)
	})
	return err
}
