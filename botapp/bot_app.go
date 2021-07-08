package botapp

import (
	"github.com/jmoiron/sqlx"
	bot "gopkg.in/tucnak/telebot.v2"
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
