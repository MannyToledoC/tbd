package connections

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect() *bun.DB{
	dsn := "postgres://postgres:postgres@host.docker.internal:5432/db?sslmode=disable"
	postgres := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn))) 

	db := bun.NewDB(postgres, pgdialect.New())
	return db
}