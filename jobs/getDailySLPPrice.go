package jobs

import (
	"time"

	"github.com/stanleygarbo/08_cron.axiemanagers.io/db"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/entities"
	"github.com/stanleygarbo/08_cron.axiemanagers.io/util"
)


func GetDailySLPPrice() {
	respData := make(map[string]map[string]float64)
	SLPPrice := entities.SLPPrice{}

	url := "https://api.coingecko.com/api/v3/simple/token_price/ethereum?contract_addresses=0xcc8fa225d80b9c7d42f96e9570156c65d6caaa25&vs_currencies=php"
	util.GetResponse(url, &respData)

	SLPPrice.Price = respData["0xcc8fa225d80b9c7d42f96e9570156c65d6caaa25"]["php"]
	SLPPrice.Date = int(time.Now().Unix())

	db.DBConn.Create(&SLPPrice)
}