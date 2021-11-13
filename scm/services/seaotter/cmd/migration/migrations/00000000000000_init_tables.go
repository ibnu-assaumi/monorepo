package migrations

var (
	migrateTables []interface{}
)

// GetMigrateTables get migrate table list
func GetMigrateTables() []interface{} {
	return migrateTables
}
