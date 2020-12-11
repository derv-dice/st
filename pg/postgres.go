package pg

import (
	"fmt"
	"github.com/jackc/pgx"
)

const PgTimeLayout = "02-01-06 15:04:05"
const PgTimeLayoutShort = "02-01-06"

// PgConnPool - returns new Postgres connection pool
func PgConnPool(login string, password string, host string, port int, name string) (pool *pgx.ConnPool, err error) {
	var cfg pgx.ConnConfig
	if cfg, err = pgx.ParseConnectionString(fmt.Sprintf("postgres://%s:%s@%s:%d/%s", login, password, host, port, name)); err != nil {
		return
	}

	pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{MaxConnections: 2, ConnConfig: cfg})
	return
}
