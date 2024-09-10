package postgres

// Extend sql.Conn
type Conn struct {
	*sqlConn
	Schema string
}

// Begin transaction
func (db *Conn) Begin(ctx Context) error {
	_, err := db.ExecContext(ctx, "BEGIN;")
	return err
}

// Commit transaction
func (db *Conn) Commit(ctx Context) error {
	_, err := db.ExecContext(ctx, "COMMIT;")
	return err
}

// Rollback transaction
func (db *Conn) Rollback(ctx Context) error {
	_, err := db.ExecContext(ctx, "ROLLBACK;")
	return err
}
