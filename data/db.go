package data

import (
	"context"
	"database/sql"
	"time"
)

func OpenDB() (*sql.DB, error) {
	dsn := "file:telemetry.db?cache=shared&_foreign_keys=on&_journal_mode=WAL&busy_timeout=5000"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.setMaxOpenConns(1)
	db.setMaxIdleConns(1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db, db.PingContext(ctx)
}

func MigrateDB(db *sql.DB) error {
	// Placeholder for migration logic
	return nil, nil
}
