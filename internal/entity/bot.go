package entity

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
	*tgbotapi.BotAPI
}

func NewTelegramBot(token string, linkHook string) (TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return TelegramBot{}, err
	}
	bot.Debug = false
	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(linkHook)); err != nil {
		return TelegramBot{}, err
	}
	return TelegramBot{bot}, nil
}
