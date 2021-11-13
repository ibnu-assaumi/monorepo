package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAlterTablemasters, downAlterTablemasters)
}

func upAlterTablemasters(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	migrateTables = append(migrateTables, &domain.master{})
	return nil
}

func downAlterTablemasters(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
