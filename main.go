package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

var cfg Config

func init() {
	ReadConfigFile(&cfg, "config/config.yml")
}

func main() {
	cron := gocron.NewScheduler(time.Local)
	cron.Every(1).Day().At("09:00").Do(dailyReport)
	cron.Every(1).Monday().At("09:00").Do(weeklyReport)
	cron.StartAsync()
}
