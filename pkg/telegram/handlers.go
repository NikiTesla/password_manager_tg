package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// handleUpdates creates main cycle of getting updates
// concurrently starts handlers of commands or messages
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			go b.handleCommand(update.Message)
			continue
		}

		// in case of basic message, write welcome information
		go b.handleStart(update.Message)
	}
}

// handleCommand routes different commands to different methods
func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case commandStart:
		if err := b.handleStart(message); err != nil {
			log.Printf("Cannot handle start command because of error: %s", err.Error())
		}
	case commandSet:
		if err := b.handleSet(message); err != nil {
			log.Printf("Cannot handle set command because of error: %s", err.Error())
		}
	case commandGet:
		if err := b.handleGet(message); err != nil {
			log.Printf("Cannot handle get command because of error: %s", err.Error())
		}
	case commandDel:
		if err := b.handleDel(message); err != nil {
			log.Printf("Cannot handle del command because of error: %s", err.Error())
		}
	default:
		if err := b.handleStart(message); err != nil {
			log.Printf("Cannot handle start command because of error: %s", err.Error())
		}
	}
}
