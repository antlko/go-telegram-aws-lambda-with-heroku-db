package botapp

import (
	"github.com/jmoiron/sqlx"
	bot "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

func CreateBot(db *sqlx.DB) (*bot.Bot, error) {
	settings := bot.Settings{
		Token:       os.Getenv("API_TOKEN"),
		Synchronous: true,
		Verbose:     true,
		Poller:      &bot.LongPoller{Timeout: 10 * time.Second},
	}

	tgBot, err := bot.NewBot(settings)
	if err != nil {
		return nil, err
	}
	initHandlers(tgBot, db)
	return tgBot, nil
}

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
