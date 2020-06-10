package web

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yfedoruck/cw3guide/pkg/env"
	"path/filepath"
)

func ImagePath(filename string) string {
	return env.BasePath() + filepath.FromSlash("/data/img/"+filename+".png")
}

func IsPhoto(command string) bool {
	switch command {
	case "herbsimg", "recipestable", "glorypoint":
		return true
	}
	return false
}

func addImages(command string, msg *tgbotapi.MessageConfig) {
	switch command {
	case "guild3":
		GloryPointTable(msg)
	case "herbs":
		HerbsTables(msg)
	}
}

func HerbsTables(msg *tgbotapi.MessageConfig) {
	var row []tgbotapi.InlineKeyboardButton

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Таблица трав", "herbsimg")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Таблица рецептов", "recipestable")
	row = append(row, tgbotapi.NewInlineKeyboardRow(btn1)...)
	row = append(row, tgbotapi.NewInlineKeyboardRow(btn2)...)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	msg.ReplyMarkup = keyboard
}

func GloryPointTable(msg *tgbotapi.MessageConfig) {
	var row []tgbotapi.InlineKeyboardButton

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	btn := tgbotapi.NewInlineKeyboardButtonData("Таблица начисления глорипоинтов", "glorypoint")
	row = append(row, tgbotapi.NewInlineKeyboardRow(btn)...)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	msg.ReplyMarkup = keyboard
}
