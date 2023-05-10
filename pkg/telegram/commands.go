package telegram

import (
	"fmt"
	"strings"

	"github.com/NikiTesla/vk_telegram/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	commandSet   = "set"
	commandGet   = "get"
	commandDel   = "del"
)

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

func (b *Bot) handleSet(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Error occured")
	defer func() {
		b.bot.Send(msg)
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
	msg.Text = fmt.Sprintf("Сохранены данные для аккаунта в %s для пользователя %s", serviceName, login)

	return nil
}

func (b *Bot) handleGet(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Error occured")
	defer func() {
		b.bot.Send(msg)
	}()

	serviceName := message.CommandArguments()
	if serviceName == "" {
		msg.Text = "Необходимо ввести имя сериса после команды:\n/get [serviceName]"
		return fmt.Errorf("not enough parameters sent. Need service name")
	}

	username, password, err := b.db.GetLoginPassword(message.From.ID, serviceName)
	if err != nil {
		msg.Text = "Произошла ошибка при поиске записи в базе данных"
		return fmt.Errorf("error occur while search for a record in database: %s", err)
	}

	msg.Text = fmt.Sprintf("Ваши данные для сервиса %s:\nusername: %s\npassword: %s", serviceName, username, password)
	return nil
}

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
