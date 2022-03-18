package mysql

import (
	"github.com/VerTox/zipcodes/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection struct {
	Connection   *gorm.DB
	Repositories *Repositories
}

type Repositories struct {
	ZipCode model.ZipCodeRepository
}

func NewConnection(dsn string) (*Connection, error) {
	conn, err := gorm.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	conn.DB().SetConnMaxLifetime(0)
	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(10)

	return &Connection{
		Connection: conn,
		Repositories: &Repositories{
			ZipCode: NewZipCodeRepository(conn),
		},
	}, nil
}

func (c *Connection) ZipCode() model.ZipCodeRepository {
	return c.Repositories.ZipCode
}

func (c *Connection) IsErrNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
