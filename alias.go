package postgres

import (
	"context"
	"database/sql"
	"time"
)

/* public alias */

type Context = context.Context
type Duration = time.Duration

/* private alias */

type sqlDB = sql.DB
type sqlConn = sql.Conn
