package config

import (
	"flag"
	"log"
)

type Conf struct {
	Addr string
	BotConf
}

type BotConf struct {
	Token   string
	Debug   bool
	TimeOut int
	Offset  int
}

func NewConf() *Conf {
	cf := Conf{}
	flag.StringVar(&cf.Token, "token", "", "telegram token")
	flag.StringVar(&cf.Addr, "addr", "localhost:50051", "the address to connect to")
	flag.Parse()
	if cf.Token == "" {
		log.Fatalln("token is empty")
	}
	if cf.Addr == "" {
		log.Fatalln("addr is empty")
	}
	cf.Debug = true
	cf.TimeOut = 60
	cf.Offset = 0
	return &cf
}
