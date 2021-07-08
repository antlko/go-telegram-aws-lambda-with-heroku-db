package botapp

import (
	"github.com/jmoiron/sqlx"
	bot "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func initHandlers(tgBot *bot.Bot, db *sqlx.DB) {
	tgBot.Handle(bot.OnText, func(m *bot.Message) {
		message := m.Text
		_, err := tgBot.Send(m.Sender, message)
		if err != nil {
			log.Println(err)
		}
	})

	tgBot.Handle("/time", func(m *bot.Message) {
		var res time.Time
		err := db.DB.QueryRow("SELECT NOW() as text").Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = tgBot.Send(m.Sender, res.String())
		if err != nil {
			log.Println(err)
		}
	})
}
