package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/mskydream/aport/config"
)

type DB struct {
	Conn *sqlx.DB
}

func (d *DB) InitDatabase(c *config.Config) *DB {
	var err error
	if d.Conn, err = sqlx.Connect("pgx", c.DatabaseSource); err != nil {
		panic(err)
	}

	m, err := migrate.New("file://db/schema", c.DatabaseSource)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if err.Error() != "no change" {
			panic(err)
		}
	}

	return d
}
