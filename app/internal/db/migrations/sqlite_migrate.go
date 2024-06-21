package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up, Up)
}

func Up(tx *sql.Tx) error {
	return nil
}

func Down(tx *sql.Tx) error {
	return nil
}
