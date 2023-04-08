package config

import "log"

type BotConf struct {
	Token   string
	Debug   bool
	TimeOut int
	Offset  int
}

func NewBotConf(token string) *BotConf {
	if token == "" {
		log.Fatalln("token is empty")
	}
	return &BotConf{
		Token:   token,
		Debug:   true,
		TimeOut: 60,
		Offset:  0,
	}
}
