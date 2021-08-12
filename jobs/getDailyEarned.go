package jobs

import (
	"fmt"
	"time"

	"github.com/stanleygarbo/08_cron.axiemanagers.io/db"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/entities"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/util"
)

func GetDailyEarned() {
	roninAddresses := []entities.RoninAddresses{}
	db.DBConn.Select([]string{"ronin", "last_read"}).Find(&roninAddresses)

	for _, address := range roninAddresses{
		if address.Ronin != ""{
			fmt.Printf("@jobs - fetching: %v", address.Ronin)
			url := fmt.Sprintf("https://game-api.skymavis.com/game-api/clients/%v/items/1", address.Ronin)
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
		}
	}
}