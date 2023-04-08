package bot

import (
	"log"
	"tgbot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
)

type BotCli struct {
	Bot     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
	GrpcCli *grpc.ClientConn
}

func NewBot(conf *config.BotConf) *BotCli {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Fatalf("new bot: %s", err.Error())
	}
	bot.Debug = conf.Debug

	u := tgbotapi.NewUpdate(conf.Offset)
	u.Timeout = conf.TimeOut

	log.Printf("authorized as %s", bot.Self.UserName)

	return &BotCli{
		Bot:     bot,
		Updates: bot.GetUpdatesChan(u),
	}
}

func (b *BotCli) Listener() {
	for update := range b.Updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			b.Bot.Send(msg)
		}
	}
}
