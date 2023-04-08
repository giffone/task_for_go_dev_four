package bot

import (
	"log"
	"tgbot/internal/config"
	"tgbot/internal/grpc_api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotCli struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
	grpcCli *grpc_api.Client
}

func NewBot(conf *config.BotConf, grpcCli *grpc_api.Client) *BotCli {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Fatalf("new bot: %s", err.Error())
	}
	bot.Debug = conf.Debug

	u := tgbotapi.NewUpdate(conf.Offset)
	u.Timeout = conf.TimeOut

	log.Printf("authorized as %s", bot.Self.UserName)

	return &BotCli{
		bot:     bot,
		updates: bot.GetUpdatesChan(u),
		grpcCli: grpcCli,
	}
}

func (b *BotCli) Listener() {
	for update := range b.updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			b.bot.Send(msg)
		}
	}
}
