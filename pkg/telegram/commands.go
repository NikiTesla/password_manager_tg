package telegram

import (
	"fmt"
	"strings"
	"time"

	"github.com/NikiTesla/vk_telegram/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	commandSet   = "set"
	commandGet   = "get"
	commandDel   = "del"
)

// handleStart handles /start route, sends welcome info message to user
func (b *Bot) handleStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, `Добро пожаловать в менеджер паролей!
	Вы можете хранить здесь свои данные для авторизации в различных сервисах
  
	Доступные команды:
	1. /set [serviceName] [username] [password]
	2. /get [serviceName]
	3. /del [serviceName]`)
	_, err := b.bot.Send(msg)

	return err
}

// handleSet is a handle function to save account data
// deleting message sent by user for not to storing password explicitly
// requires three arguments from user to be sent with command /set
func (b *Bot) handleSet(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Error occured")
	defer func() {
		b.bot.Send(msg)
	}()

	// deletion of message with password after 30 seconds
	go func() {
		time.Sleep(30 * time.Second)
		deleteConfig := tgbotapi.DeleteMessageConfig{
			ChatID:    message.Chat.ID,
			MessageID: message.MessageID,
		}
		b.bot.DeleteMessage(deleteConfig)
	}()

	args := strings.Split(message.CommandArguments(), " ")
	if len(args) < 3 {
		msg.Text = "Необходимо ввести имя сериса, имя пользователя и пароль после команды:\n/set [serviceName] [username] [password]"
		return fmt.Errorf("not enough parameters sent. Need service name, login and password")
	}

	serviceName, login, password := args[0], args[1], args[2]

	err := b.db.CreateLoginPassword(&repository.LoginPassword{
		UserID:      int(message.Chat.ID),
		ServiceName: serviceName,
		Login:       login,
		Password:    password,
	})

	if err != nil {
		msg.Text = "Произошла ошибка при записи в базу данных"
		return fmt.Errorf("error occur while creating a record in database: %s", err)
	}
	msg.Text = fmt.Sprintf("Сохранены данные для аккаунта в %s", serviceName)

	return nil
}

// handleGet is a handle function to get saved account data
// deleting message sent by bot for not to storing password explicitly
// requires one argument from user to be sent with command /get
func (b *Bot) handleGet(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Error occured")

	serviceName := message.CommandArguments()
	if serviceName == "" {
		msg.Text = "Необходимо ввести имя сериса после команды:\n/get [serviceName]"
		b.bot.Send(msg)
		return fmt.Errorf("not enough parameters sent. Need service name")
	}

	username, password, err := b.db.GetLoginPassword(message.From.ID, serviceName)
	if err != nil {
		msg.Text = "Произошла ошибка при поиске записи в базе данных"
		b.bot.Send(msg)
		return fmt.Errorf("error occur while search for a record in database: %s", err)
	}

	msg.Text = fmt.Sprintf("Ваши данные для сервиса %s:\nusername: %s\npassword: %s", serviceName, username, password)
	sent_message, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("cannot send message with text %s", msg.Text)
	}

	// deletion of message with password after 30 seconds
	go func() {
		time.Sleep(30 * time.Second)
		deleteConfig := tgbotapi.DeleteMessageConfig{
			ChatID:    sent_message.Chat.ID,
			MessageID: sent_message.MessageID,
		}
		b.bot.DeleteMessage(deleteConfig)
	}()

	return nil
}

// handleDel is a handle function to delete saved account data
// requires one argument from user to be sent with command /del
func (b *Bot) handleDel(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Error occured")
	defer func() {
		b.bot.Send(msg)
	}()

	serviceName := message.CommandArguments()
	if serviceName == "" {
		msg.Text = "Необходимо ввести имя сериса после команды:\n/del [serviceName]"
		return fmt.Errorf("not enough parameters sent. Need service name, username")
	}

	err := b.db.DeleteLoginPassword(int(message.Chat.ID), serviceName)
	if err != nil {
		msg.Text = "Произошла ошибка при удалении записи из базы данных"
		return fmt.Errorf("error occur while deleting a record in database: %s", err)
	}
	msg.Text = fmt.Sprintf("Данные для сервиса %s удалены", serviceName)

	return nil
}
