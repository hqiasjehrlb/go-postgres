package postgres

import (
	"fmt"
)

type DataSource struct {
	Host        string // default: localhost
	Port        uint16 // default: 5432
	User        string // default: postgres
	Password    string
	DBName      string
	SSLMode     SSLMode  // default: SSLModeDisable
	ConnTimeout Duration // default: 10 sec
}

// Return dataSourceName string
func (opt *DataSource) String() string {
	var (
		host     = opt.Host
		port     = opt.Port
		user     = opt.User
		password = opt.Password
		dbname   = opt.DBName
		sslmode  = opt.SSLMode
		timeout  = uint32(opt.ConnTimeout.Seconds())
	)

	if host == "" {
		host = "localhost"
	}
	if port == 0 {
		port = 5432
	}
	if user == "" {
		user = "postgres"
	}
	if timeout <= 0 {
		timeout = 10
	}

	return fmt.Sprintf(
		"host=%s port=%d sslmode=%s user=%s password=%s dbname=%s connect_timeout=%d",
		host, port, sslmode.String(), user, password, dbname, timeout,
	)
}
