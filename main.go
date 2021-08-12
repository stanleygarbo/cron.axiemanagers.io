package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/db"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/jobs"
)

func main() {
	db.InitDatabase()

	c := cron.New()
	c.AddFunc("@every 1h30m", jobs.GetDailySLPPrice)
	c.AddFunc("@daily", jobs.GetDailyEarned)
	c.Start()

	c.Stop()
	fmt.Print("terminated")
}