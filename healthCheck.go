package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
)

//GetHealthCheck retrieve from given URL the Service Status
func GetHealthCheck() (body []byte, err error) {
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

//SendHealthCheck send to Telegram bot the retrieved data
func SendHealthCheck(bot *tgbotapi.BotAPI, chatID int64) {
	message := tgbotapi.NewMessage(chatID, "What can do this bot?")

	body, err := GetHealthCheck()
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