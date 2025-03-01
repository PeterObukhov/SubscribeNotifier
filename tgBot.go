package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

var isSubscriptionInput bool
var tgBot *bot.Bot
var ctx context.Context

func startBot() {
	var cancel context.CancelFunc
	ctx, cancel = signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	var err error
	tgBot, err = bot.New(configuration.Token, opts...)

	if err != nil {
		panic(err)
	}

	tgBot.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := inline.New(b).
		Row().
		Button("Добавить подписку", []byte("newSubs"), onInlineKeyboardSelect)

	if isSubscriptionInput {
		splitMsg := strings.Split(update.Message.Text, ":")
		price, _ := strconv.ParseFloat(strings.Trim(splitMsg[1], " "), 32)
		date, _ := strconv.ParseInt(strings.Trim(splitMsg[2], " "), 10, 32)
		subscription := Subscription{Name: splitMsg[0], Price: float32(price), Date: int(date)}
		user := User{Login: update.Message.From.Username, ChatId: update.Message.Chat.ID, Subscriptions: []Subscription{subscription}}

		if addUserOrSubscription(*db, user) {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      update.Message.Chat.ID,
				Text:        "Информацию о подписке добавлена",
				ReplyMarkup: kb,
			})
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      update.Message.Chat.ID,
				Text:        "Ошибка, информацию о подписке не добавлена",
				ReplyMarkup: kb,
			})
		}

		isSubscriptionInput = false
	} else {
		if update.Message != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      update.Message.Chat.ID,
				Text:        "Добавить подписку?",
				ReplyMarkup: kb,
			})
		}
	}
}

func onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: mes.Message.Chat.ID,
		Text:   "Введите данные о новой подписке в формате *Название* : *Сумма* : *Число, в которое происходит списание*",
	})

	isSubscriptionInput = true
}
