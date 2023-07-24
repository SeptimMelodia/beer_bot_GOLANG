package main

import (
	"log"
	"tg_bot_beer_to_peer/pkg/repository"
	"tg_bot_beer_to_peer/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(viper.GetString("tg.token"))
	if err != nil {
		log.Panic(err)
	}

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
	rep := repository.NewRepository(db)
	user, err := rep.Authorization.GetUser("505780494")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("User is name %s found.\n", user.Name)

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../../configs") // Для дебага в VS Cpde
	return viper.ReadInConfig()
}
