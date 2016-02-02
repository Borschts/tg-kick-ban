package main

import (
	"github.com/codegangsta/cli"
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
)

func tgmigrate(c cli.Context) {
	allErrors, ok := migrate.UpSync("driver://url", "./path")
	println(allErrors, ok)
}
