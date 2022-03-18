package repository

import (
	"github.com/VerTox/zipcodes/domain"
	"github.com/VerTox/zipcodes/domain/model"
	"github.com/VerTox/zipcodes/repository/mysql"
)

type Connection struct {
	mysql *mysql.Connection
}

func New(m *mysql.Connection) domain.Connection {
	return &Connection{
		mysql: m,
	}
}

func (c *Connection) ZipCode() model.ZipCodeRepository {
	return c.mysql.ZipCode()
}

func (c *Connection) IsErrNotFound(err error) bool {
	return c.mysql.IsErrNotFound(err)
}
