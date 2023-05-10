package telegram

import (
	"fmt"
	"log"

	"github.com/NikiTesla/vk_telegram/pkg/environment"
	"github.com/NikiTesla/vk_telegram/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	env *environment.Environment
	bot *tgbotapi.BotAPI
	db  repository.DataBase
}

// NewBot creates NewBotApi, returns structure with it, environment and database as Bot
func NewBot(env *environment.Environment, botToken string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("can't create Bot API, error: %s", err.Error())
	}
	bot.Debug = env.Config.Debug

	return &Bot{
		env: env,
		bot: bot,
		db:  &repository.PostgresDB{DB: env.DB},
	}, nil
}

// Start inititates updates channel and starts listening for updates
func (b *Bot) Start() error {
	log.Printf("Bot authorized on account %s is running", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)
	return nil
}

// initUpdatesChannel initiates channel for updates
func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
