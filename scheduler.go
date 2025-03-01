package main

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/go-telegram/bot"
	"github.com/google/uuid"
)

var scheduler gocron.Scheduler

func startScheduler() {
	var err error
	scheduler, err = gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	scheduler.Start()
}

func addJob(date int, subscription string, chatId int64) {
	_, err := scheduler.NewJob(
		gocron.MonthlyJob(
			1,
			gocron.NewDaysOfTheMonth(date),
			gocron.NewAtTimes(
				gocron.NewAtTime(12, 0, 0),
			),
		),
		gocron.NewTask(
			func(supscriptionName string, chatId int64) {
				tgBot.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatId,
					Text:   supscriptionName,
				})
			},
			subscription,
			chatId,
		),
	)
	if err != nil {
		panic(err)
	}
}

func removeJob(id uuid.UUID) {
	scheduler.RemoveJob(id)
}
