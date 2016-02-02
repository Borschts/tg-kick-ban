package main

import (
	"github.com/codegangsta/cli"
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
	"os"
)

func tgmigrate(c *cli.Context) {
	allErrors, ok := migrate.UpSync(c.String("pgconn"), "./db")
	if !ok {
		println("tgmigrate errors:")
		for _, err := range allErrors {
			println("\t" + err.Error())
		}
		os.Exit(DBMIGRATION_ERROR)
	}
}
