package main

import (
	"github.com/codegangsta/cli"
	"github.com/tucnak/telebot"
	"os"
	"time"
)

var theBot *telebot.Bot

func newbot(c *cli.Context) {
  var err error
	theBot, err = telebot.NewBot(c.String("tgtoken"))
	if err != nil {
		println(err.Error())
		os.Exit(TG_ERROR)
	}
	theBot.Messages = make(chan telebot.Message, 1000)
	theBot.Queries = make(chan telebot.Query, 1000)
	go messages()
	go queries()
	theBot.Start(1 * time.Second)
}

func messages() {
	for message := range theBot.Messages {
		if message.Chat.IsGroupChat() {
			println("got group")
		} else {
			theBot.SendMessage(message.Chat, getHelp(), nil)
		}
	}
}

func queries() {
	for range theBot.Queries {
    println("got query")
  }
}
