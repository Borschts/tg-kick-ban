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
}
