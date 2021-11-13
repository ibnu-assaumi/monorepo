package migrations

import (
	"database/sql"
	"monorepo/services/seaotter/pkg/shared/domain"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAlterTableSalesorders, downAlterTableSalesorders)
}

func upAlterTableSalesorders(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	migrateTables = append(migrateTables, &domain.Salesorder{})
	return nil
}

func downAlterTableSalesorders(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}	
