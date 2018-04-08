package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Persistor interface {
	Add(object interface{}) (int64, error)
	Update(object interface{})
	Delete(object interface{})
}

type generalPersistor struct {
	connection *xorm.Engine
}

func NewGeneralPersistor(connectionName string) (Persistor, error) {
	c, err := GetConnectionManager().GetConnection(connectionName)

	p := &generalPersistor{
		connection: c.adapter,
	}

	return p, err
}

func (p *generalPersistor) Add(o interface{}) (int64, error) {
	return p.connection.Insert(o)
}

func (p *generalPersistor) Update(o interface{}) {

}

func (p *generalPersistor) Delete(o interface{}) {

}
