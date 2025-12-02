package migrate

import (
	"context"
	"database/sql"
	"embed"
	"strings"
)

var files embed.FS

func Run(ctx context.Context, db *sql.DB) error {
	entries, err := files.ReadDir("migrations")

	if err != nil {
		return err
	}

	for _, entry := range entries {
		b, _ := files.ReadFile("../migrations/" + e.Name())
		stmts := strings.Split(string(b), ";")
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		for _, s := range stmts {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if _, err := tx.ExecContext(ctx, s); err != nil {
				_ = tx.Rollback()
				return err
			}
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}
