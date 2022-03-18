package main

import (
	"fmt"
	"github.com/VerTox/zipcodes/repository"
	"github.com/VerTox/zipcodes/repository/mysql"
	"github.com/davecgh/go-spew/spew"
	"os"
)

func main() {
	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		panic("No APP_PORT provided")
	}

	dsn := os.Getenv("MYSQL_DSN")

	if dsn == "" {
		panic("No MYSQL_DSN provided")
	}

	m, err := mysql.NewConnection(dsn)

	if err != nil {
		panic(err)
	}

	c := repository.New(m)

	a := &Application{
		Connection: c,
	}
	spew.Dump("started")
	a.Run(fmt.Sprintf(":%s", appPort))
}
