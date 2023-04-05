package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//InitChat keeps bot refreshed every 50s and provide with Telegram Bot commands
func InitChat(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		fmt.Println(update.Message.Chat.ID)
		fmt.Println(update.Message.Text)

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "What can do this bot?")
		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.Text = "Hello New User!I understand /help and /status."
		case "help":
			msg.Text = "Hi :)"
		case "status":
			body, err := GetHealthCheck()
			if err != nil {
				fmt.Println(err)
				return
			}
			msg.Text = PrettyString(string(body))
		default:
			msg.Text = "I don't know that command"
		}
		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}
