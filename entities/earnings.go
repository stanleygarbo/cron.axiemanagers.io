package entities

import "gorm.io/gorm"

type Earning struct {
	gorm.Model
	Ronin  string
	Date   int
	Earned int
}
