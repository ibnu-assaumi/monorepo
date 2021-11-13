// Code generated by candi v1.8.8.

package main

import (
	"flag"
	"log"
	"os"

	"monorepo/services/seaotter/cmd/migration/migrations"

	"github.com/Bhinneka/candi/config/database"
	"github.com/Bhinneka/candi/config/env"

	"github.com/Bhinneka/candi/candihelper"

	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	env.Load("seaotter")
	sqlDeps := database.InitSQLDatabase()

	flags.Parse(os.Args[1:])
	args := flags.Args()
	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	dir := os.Getenv("WORKDIR") + "cmd/migration/migrations"
	switch args[0] {
	case "create":
		migrationType := "sql"
		if len(args) > 2 && args[2] == "init_table" {
			migrationType = "go"
		}
		if err := goose.Create(sqlDeps.WriteDB(), dir, args[1], migrationType); err != nil {
			log.Fatalf("goose %v: %v", args[1], err)
		}

	default:

		if err := goose.Run(args[0], sqlDeps.WriteDB(), dir, arguments...); err != nil {
			log.Fatalf("goose %v: %v", args[0], err)
		}

		if migrateTables := migrations.GetMigrateTables(); len(migrateTables) > 0 {
			gormWrite, err := gorm.Open(postgres.New(postgres.Config{
				Conn: sqlDeps.WriteDB(),
			}), &gorm.Config{
				SkipDefaultTransaction:                   true,
				DisableForeignKeyConstraintWhenMigrating: true,
			})
			if err != nil {
				log.Fatal(err)
			}
			tx := gormWrite.Begin()
			if err := gormWrite.AutoMigrate(migrateTables...); err != nil {
				tx.Rollback()
				log.Fatal(err)
			}
			tx.Commit()
		}
	}
	log.Printf("\x1b[32;1mMigration to \"%s\" success\x1b[0m\n", candihelper.MaskingPasswordURL(env.BaseEnv().DbSQLWriteDSN))
}
