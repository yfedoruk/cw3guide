package web

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yfedoruck/cw3guide/pkg/env"
	"github.com/yfedoruck/cw3guide/pkg/fail"
	"os"
	"path/filepath"
)

type config struct {
	TelegramBotToken string
}

func Token() string {
	file, err := os.Open(env.BasePath() + filepath.FromSlash("/config.json"))
	fail.Check(err)

	decoder := json.NewDecoder(file)
	configuration := config{}

	err = decoder.Decode(&configuration)
	fail.Check(err)

	return configuration.TelegramBotToken
}

func Updates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	switch os.Getenv("USERDOMAIN") {
	case "localhost", "home":
		return longPooling(bot)
	default:
		return webHooks(bot)
	}
}

// long pooling for localhost
func longPooling(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	_, err := bot.RemoveWebhook()
	fail.Check(err)

	updates, err := bot.GetUpdatesChan(u)
	fail.Check(err)

	return updates
}

// web hooks for awake heroku from idling
func webHooks(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {

	conf := tgbotapi.NewWebhook("https://cw3guide.herokuapp.com/" + bot.Token)
	_, err := bot.SetWebhook(conf)
	fail.Check(err)

	return bot.ListenForWebhook("/" + bot.Token)
}
