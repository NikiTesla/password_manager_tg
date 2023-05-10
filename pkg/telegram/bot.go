package telegram

import (
	"fmt"
	"log"

	"github.com/NikiTesla/vk_telegram/pkg/environment"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	env *environment.Environment
	bot *tgbotapi.BotAPI
}

func NewBot(env *environment.Environment, botToken string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("can't create Bot API, error: %s", err.Error())
	}
	bot.Debug = env.Config.Debug

	return &Bot{
		env: env,
		bot: bot,
	}, nil
}

func (b *Bot) Start() error {
	log.Printf("Bot authorized on account %s is running", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				log.Print("error occured while command handling, error: ", err)
			}
			continue
		}

		b.handleMessage(update.Message)
	}
}
