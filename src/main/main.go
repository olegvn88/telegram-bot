package main

import (
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"joke"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type JsonData struct {
	Token string `json:"token"`
}

var buttons = []tgbotapi.KeyboardButton{
	tgbotapi.KeyboardButton{Text: "Get Joke"},
}

//const WebhookURL = "https://english891.herokuapp.com/"
const s = "Welcome to the help menu of boot english891!\n - get joke \n - time \n - help \n"
const jokeUrl = "http://api.icndb.com/jokes/random?limitTo=[nerdy]"

func main() {
	// Heroku прокидывает порт для приложения в переменную окружения PORT
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI(ParseJson().Token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//// Устанавливаем вебхук
	//_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	//if err != nil {
	//	log.Fatal(err)
	//}

	//===
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	//======
	//updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	//получаем все обновления из канала updates
	for update := range updates {
		var message tgbotapi.MessageConfig
		log.Println("received text: ", update.Message.Text)

		switch strings.ToLower(update.Message.Text) {
		case "help":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, s)
		case "get joke":
			//если пользователь нажал на кнопку,то прийдет сообщение "Get Joke"
			message = tgbotapi.NewMessage(update.Message.Chat.ID, joke.GetJoke(jokeUrl))
		case "time":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, time.Now().String())
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `Press Get Joke to receive joke or type help`)
		}

		//в ответном сообщении просим показать клавиатуру

		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

		_, _ = bot.Send(message)

	}
}

func ParseJson() JsonData {
	jData, err := ioutil.ReadFile("/home/onest/go/src/github.com/olegvn88/bot/src/main/propreties.config")
	if err != nil {
		panic(err)
	}
	data := JsonData{}
	err = json.Unmarshal(jData, &data)
	if err != nil {
		panic(err)
	}
	return data
}
