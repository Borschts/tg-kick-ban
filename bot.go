package main

import (
	"github.com/codegangsta/cli"
	"github.com/tucnak/telebot"
	"os"
)

var theBot *telebot.Bot

func newbot(c *cli.Context) {
	bot, err := telebot.NewBot(c.String("tgtoken"))
	if err != nil {
		println(err.Error())
		os.Exit(TG_ERROR)
	}
	theBot = bot
}
