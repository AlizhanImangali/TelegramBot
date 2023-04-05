package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	config := tgbotapi.NewUpdate(0)
	config.Timeout = 60
	channel := bot.GetUpdatesChan(config)

	crontab := cron.New()
	crontab.AddFunc("@every 1h", func() {
		SendHealthCheck(bot, -1001594451527)
	})
	crontab.Start()
	//crontab.Stop()
	// init channel
	InitChat(bot, channel)
}
