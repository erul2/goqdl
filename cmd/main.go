package main

import (
	"github.com/erul2/goqdl"
	"github.com/erul2/goqdl/bot"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	botSettings := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b := bot.NewTeleBotServer(botSettings)

	goQDL, err := initGoQdl()
	if err != nil {
		log.Fatalln("error at init goqdl: ", err)
	}

	handler := bot.NewBotHandler(goQDL)

	startTeleBotServer(b, handler)
}

func startTeleBotServer(b bot.TeleBotServer, handler *bot.Handler) {
	b.SetupRoutes(func(b *bot.TeleBotServer) {
		b.Bot.Handle("/download", handler.DownloadAlbum)
	})
	log.Println("bot satrted...")
	b.Start()
}

func initGoQdl() (*goqdl.GoQDL, error) {
	goQDL, err := goqdl.NewGoQDL(os.Getenv("QOBUZ_EMAIL"), os.Getenv("QOBUZ_PASSWORD"), goqdl.GoQDLConfig{
		DownloadDir:     os.Getenv("QOBUZ_DOWNLOAD_DIR"),
		DownloadQuality: os.Getenv("QOBUZ_DOWNLOAD_QUALITY"),
	})
	if err != nil {
		return nil, err
	}
	return goQDL, nil
}
