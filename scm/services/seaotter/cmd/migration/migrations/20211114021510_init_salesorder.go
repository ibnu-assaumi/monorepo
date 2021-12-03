package migrations

import (
	"database/sql"
	"monorepo/services/seaotter/pkg/shared/model"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAlterTableSalesorders, downAlterTableSalesorders)
}

func upAlterTableSalesorders(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	migrateTables = append(migrateTables, &model.Salesorder{})
	return nil
}

func downAlterTableSalesorders(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
