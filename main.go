package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"net/http"
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
	crontab.AddFunc("@every 15s", func() {
		SendHealthCheck(bot, -1001594451527)
	})
	crontab.Start()

	// init channel
	InitChat(bot, channel)
}
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
			body, err := GetHealthCheckStatus()
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

func SendHealthCheck(bot *tgbotapi.BotAPI, chatID int64) {
	message := tgbotapi.NewMessage(chatID, "What can do this bot?")

	body, err := GetHealthCheckStatus()
	if err != nil {
		fmt.Println(err)
		return
	}
	message.Text = PrettyString(string(body))

	if _, err = bot.Send(message); err != nil {
		fmt.Println(err)
		return
	}

	return
}

func GetHealthCheckStatus() (body []byte, err error) {
	url := "https://library-service-9t55.onrender.com/api/health"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer HKRTJH IEJHIOJHAEJPHJ")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func PrettyString(str string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return ""
	}
	return prettyJSON.String()
}
