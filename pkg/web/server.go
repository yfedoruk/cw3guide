package web

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yfedoruck/cw3guide/pkg/env"
	"github.com/yfedoruck/cw3guide/pkg/fail"
	"log"
	"net/http"
)

type Server struct {
	Port string
}

func (s *Server) Start() {

	bot, err := tgbotapi.NewBotAPI(Token())
	fail.Check(err)

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := Updates(bot)

	log.Println("Starting web server on", s.Port)
	go func() {
		if err := http.ListenAndServe(":"+s.Port, nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()

	var fw = NewFlyweight()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		var msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
		var command = update.Message.Command()

		msg.Text = fw.GetPage(command)
		msg.ParseMode = "markdown"
		msg.DisableWebPagePreview = true

		_, err := bot.Send(msg)
		fail.Warning(err)
	}

}

func NewServer() *Server {
	s := &Server{}
	s.Port = env.Port()
	http.HandleFunc("/", MainHandler)
	return s
}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	_, err := resp.Write([]byte("Hi there! I'm Chat wars 3 guide bot!"))
	fail.Check(err)
}
