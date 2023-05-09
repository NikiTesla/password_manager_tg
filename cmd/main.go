package main

import (
	"log"
	"os"

	"github.com/NikiTesla/vk_telegram/pkg/environment"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("can't load env variables, err:", err.Error())
	}

	configFile := os.Getenv("CONFIGFILE")
	_, err := environment.NewEnvironment(configFile)
	if err != nil {
		log.Fatal("can't load environment, err:", err.Error())
	}
}
