package main

import (
	"github.com/codegangsta/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:   "tgtoken",
		Usage:  "the Telegram bot token",
		EnvVar: "TG_TOKEN",
	},
	cli.StringFlag{
		Name:   "pgconn",
		Value:  "postgres://user@host:5432/database",
		Usage:  "the Postgress connection string",
		EnvVar: "PG_CONN",
	},
}
