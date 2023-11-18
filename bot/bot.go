package bot

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
)

type TeleBotServer struct {
	Bot *tele.Bot
}

func NewTeleBotServer(settings tele.Settings) TeleBotServer {
	teleBotServer, err := tele.NewBot(settings)
	if err != nil {
		log.Fatalln("error at init bot: ", err)
	}
	return TeleBotServer{
		Bot: teleBotServer,
	}
}

func (b *TeleBotServer) SetupRoutes(fc func(b *TeleBotServer)) {
	fc(b)
}

func (b *TeleBotServer) Start() {
	b.Bot.Start()
}

func (b *TeleBotServer) SendMessage(to int64, message string) (*int, *int64, error) {
	user := tele.User{
		ID: to,
	}
	msg, err := b.Bot.Send(&user, message)
	if err != nil {
		return nil, nil, err
	}
	return &msg.ID, &msg.Chat.ID, nil
}

func (b *TeleBotServer) EditMessage(messageId int, chatId int64, updatedText string) (*int, *int64, error) {
	msg := tele.StoredMessage{
		MessageID: fmt.Sprintf("%d", messageId),
		ChatID:    chatId,
	}
	edit, err := b.Bot.Edit(&msg, updatedText)
	if err != nil {
		return nil, nil, err
	}
	return &edit.ID, &edit.Chat.ID, nil
}
