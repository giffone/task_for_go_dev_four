package main

import (
	"flag"
	"log"
	"tgbot/internal/bot"
	"tgbot/internal/config"
	"tgbot/internal/grpc_api"
)

func main() {
	var addr, token string
	flag.StringVar(&token, "token", "", "telegram token")
	flag.StringVar(&addr, "addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	// grpc cli
	grpcCli, err := grpc_api.NewGrpcClient(addr)
	if err != nil {
		log.Fatalf("grpc client: %v", err)
	}
	defer grpcCli.Close()

	// configuration
	conf := config.NewBotConf(token)

	// telegram cli
	b := bot.NewBot(conf, grpcCli)
	b.Listener()
}
