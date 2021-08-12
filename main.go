package main

import (
	"fmt"
	"sync"

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

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
	fmt.Print("terminated")
}