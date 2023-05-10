package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	commandStart = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStart(message)
	default:
		return b.handleUnknown(message)
	}
}

func (b *Bot) handleStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to personal password manager!")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleUnknown(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command "+message.Text)

	_, err := b.bot.Send(msg)
	return err
}
