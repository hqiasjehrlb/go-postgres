package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Initialize new DB pool by config
func OpenByConf(config DataSource) (*Pool, error) {
	return Open(config.String())
}

// Initialize new DB pool by environment variables
//
// → https://www.postgresql.org/docs/9.3/libpq-envars.html
func OpenDefault() (*Pool, error) {
	return Open("")
}

// Initialize new DB pool by dataSourceName string
func Open(dataSourceName string) (*Pool, error) {
	p, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		if p != nil {
			p.Close()
		}
		return nil, err
	}
	return &Pool{p}, nil
}

// Extend sql.DB
type Pool struct {
	*sqlDB
}

// Acquire a new session
func (pool *Pool) Conn(ctx Context) (*Conn, error) {
	conn, err := pool.sqlDB.Conn(ctx)
	if err != nil {
		if conn != nil {
			conn.Close()
		}
		return nil, err
	}
	return &Conn{conn, `"public"`}, nil
}
