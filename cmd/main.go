package main

import (
	"log"
	"os"

	"github.com/NikiTesla/vk_telegram/pkg/environment"
	"github.com/NikiTesla/vk_telegram/pkg/telegram"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("can't load env variables, err:", err.Error())
	}

	configFile := os.Getenv("CONFIGFILE")
	env, err := environment.NewEnvironment(configFile)
	if err != nil {
		log.Fatal("can't load environment, err:", err.Error())
	}

	bot, err := telegram.NewBot(env, os.Getenv("TELEBOT_API"))
	if err != nil {
		log.Fatal("Error occured while creating Bot API: ", err.Error())
	}

	if err := bot.Start(); err != nil {
		log.Fatal("Error occured while running bot, error: ", err)
	}

}
