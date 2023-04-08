package config

import (
	"flag"
	"log"
)

type Conf struct {
	Addr string
	Db   DbConf
}

type DbConf struct {
	Driver string
	Conn   string
}

func NewConf() *Conf {
	cf := Conf{}
	flag.StringVar(&cf.Db.Driver, "driver", "postgres", "database driver name")
	flag.StringVar(&cf.Db.Conn, "connection", "", "connection to database")
	flag.StringVar(&cf.Addr, "addr", "localhost:50051", "the address server")
	flag.Parse()
	if cf.Addr == "" {
		log.Fatalln("conf: addr is empty")
	}
	if cf.Db.Driver == "" {
		log.Fatalln("conf: driver name is empty")
	}
	if cf.Db.Conn == "" {
		log.Fatalln("conf: connection is empty")
	}
	return &cf
}
