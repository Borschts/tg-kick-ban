package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tg-kick-ban"
	app.Usage = "kick, bans and purge user messages in Telegram groups"
	app.Flags = flags
	app.Action = newbot
	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
		os.Exit(CLI_ERROR)
	}
}
