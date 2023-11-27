package models

import "time"

type Commodity struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	DefaultName    string    `json:"default_name" gorm:"not null"`
	ProduceAt      time.Time `json:"produce_at" gorm:"not null"`
	ProduceAddress string    `json:"produce_address" gorm:"not null"`
	Category       string    `json:"category" gorm:"not null"`
}
