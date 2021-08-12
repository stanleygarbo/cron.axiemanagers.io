package entities

import "gorm.io/gorm"

type SLPPrice struct {
	gorm.Model
	Price float64
	Date  int
}