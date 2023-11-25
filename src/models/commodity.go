package models

import "time"

type Commodity struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	DefaultName    string    `json:"default_name"`
	ProduceAt      time.Time `json:"produce_at"`
	ProduceAddress string    `json:"produce_address"`
}
