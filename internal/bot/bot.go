package bot

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)
// New creates a new instance of a Telegram bot.
func New(path string) *tele.Bot{
	if err := godotenv.Load(path); err!= nil {
		panic("Error loading.env file")
    }
	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		panic("error creating bot")
	}
	return b
}