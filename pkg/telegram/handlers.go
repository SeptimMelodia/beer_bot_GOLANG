package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const slashCommandStart = "start"

func (b *Bot) handleMassage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	//msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	//msg.ReplyToMessageID = message.MessageID

	//b.bot.Send(msg)
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command")
	switch message.Command() {
	case slashCommandStart:
		msg.Text = "Start commnd"
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
