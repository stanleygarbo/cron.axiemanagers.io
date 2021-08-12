package entities

import (
	"time"

	"gorm.io/gorm"
)

type RoninAddresses struct {
	gorm.Model
	Ronin    string    `gorm:"primaryKey"`
	LastRead time.Time `json:"last_read"`
}
