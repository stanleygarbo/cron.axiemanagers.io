package jobs

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/stanleygarbo/08_cron.axiemanagers.io/db"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/entities"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/util"
	"golang.org/x/sync/semaphore"
)

func GetDailyEarned() {
	roninAddresses := []entities.RoninAddresses{}
	db.DBConn.Select([]string{"ronin", "last_read"}).Where("last_read >= NOW() - INTERVAL '48 HOURS'").Find(&roninAddresses)

	wg := &sync.WaitGroup{}
	sem := semaphore.NewWeighted(3)

	for i, address := range roninAddresses{
		if address.Ronin != ""{
			wg.Add(1)
			if err := sem.Acquire(context.Background(), 1); err != nil{
				log.Fatalf("Error: %v\n", err)
			}

			go func(i int,ronin string){
				// fmt.Printf("@jobs - fetching: %v", address.Ronin)
				defer wg.Done()
				defer sem.Release(1)
				url := fmt.Sprintf("https://game-api.skymavis.com/game-api/clients/%v/items/1", ronin)
				lunaciaResponse := entities.LunaciaResponse{}
				err := util.GetResponse(url, &lunaciaResponse)
				if err != nil{
					fmt.Printf("@jobs - get daily earned::::Error fetching response:%v\n", err)
				} else {
					earning := entities.Earning{
						Ronin: lunaciaResponse.RoninAddress,
						Date: int(time.Now().Unix()),
						Earned: lunaciaResponse.TotalSLP,
					}

					db.DBConn.Create(&earning)
				}
			}(i, address.Ronin)
		}
	}

	wg.Wait()
}