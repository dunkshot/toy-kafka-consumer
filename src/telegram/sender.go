package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// token : 5290721096:AAFWLl_U0IDF1MFSbjtZtwQAOJ1CkbuMtW8
// chat id : 370311245

func Sender(txt string) {
	bot, err := tgbotapi.NewBotAPI("5290721096:AAFWLl_U0IDF1MFSbjtZtwQAOJ1CkbuMtW8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(370311245, txt)
	bot.Send(msg)

}
