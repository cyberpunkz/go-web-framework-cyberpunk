package persistence

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var cm *ConnectionManager
var once sync.Once

type ConnectionManager struct {
	connections map[string]*Connection
}

type Connection struct {
	adapter *xorm.Engine
}

func GetConnectionManager() *ConnectionManager {
	once.Do(func() {
		cm = &ConnectionManager{
			connections: make(map[string]*Connection),
		}
	})

	return cm
}

func (cm *ConnectionManager) AddConnection(
	name string,
	driver string,
	host string,
	port string,
	database string,
	user string,
	password string,
) (err error) {

	if cm.connections[name] != nil {
		return fmt.Errorf(
			"connection with the same name already exists: %s",
			name,
		)
	}

	cm.connections[name], err = newConnection(
		driver,
		host,
		port,
		database,
		user,
		password,
	)

	return err
}

func (cm *ConnectionManager) GetConnection(name string) (*Connection, error) {
	var err error

	c := cm.connections[name]

	if c == nil {
		err = fmt.Errorf(
			"connection doesn't exist with name: %s",
			name,
		)
	}

	return c, err
}

func newConnection(
	driver string,
	host string,
	port string,
	database string,
	user string,
	password string,
) (*Connection, error) {
	adapter, err := xorm.NewEngine(
		driver,
		user+":"+password+"@("+host+":"+port+")/"+database,
	)

	c := &Connection{
		adapter: adapter,
	}

	return c, err
}
