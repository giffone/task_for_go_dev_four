package main

import (
	"flag"
	"tgbot/internal/bot"
	"tgbot/internal/config"
)

func main() {
	var token string
	flag.StringVar(&token, "token", "", "telegram token")
	flag.Parse()

	// configuration
	conf := config.NewBotConf(token)

	// grpc cli


	// telegram cli
	b := bot.NewBot(conf)
}
