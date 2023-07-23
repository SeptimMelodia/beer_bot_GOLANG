package main

import (
	"log"
	"tg_bot_beer_to_peer/pkg/repository"
	"tg_bot_beer_to_peer/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5837562265:AAFJYENJh2I0QzZDO9hs2YBKIkzOG-vqBrA")
	if err != nil {
		log.Panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	log.Println(viper.GetString("db.host"))
	db, err := repository.NewMySqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatalf("failed to initiaslize db: %s", err.Error())
	}
	repository.NewRepository(db)

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
