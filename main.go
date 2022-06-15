package main

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/cache"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/db"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/jobs"
)

func main() {
	db.InitDatabase()
	cache.InitRedis()

	c := cron.New()
	c.AddFunc("@daily", jobs.GetDailyEarned)
	c.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
	fmt.Print("terminated")
}