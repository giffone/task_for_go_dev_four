package main

import (
	"log"
	"tgbot/internal/bot"
	"tgbot/internal/config"
	"tgbot/internal/grpc_api"
)

func main() {
	// configuration
	conf := config.NewConf()

	// grpc cli
	grpcCli, err := grpc_api.NewGrpcClient(conf.Addr)
	if err != nil {
		log.Fatalf("grpc client: %v", err)
	}
	defer grpcCli.Close()

	// telegram cli
	b := bot.NewBot(&conf.BotConf, grpcCli)
	b.Listener()
}
