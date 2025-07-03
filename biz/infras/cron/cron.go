package cron

import (
	"time"
)

func InitCron() {

	go run()

}
func run() {

	tickerSecond_5 := time.NewTicker(5 * time.Second)
	tickerSecond_30 := time.NewTicker(30 * time.Second)
	tickerHour_6 := time.NewTicker(6 * time.Hour)

	for {
		select {
		case <-tickerSecond_5.C:

		case <-tickerSecond_30.C:

		case <-tickerHour_6.C:

			setResponseAnswersCount()
		}

	}
}
