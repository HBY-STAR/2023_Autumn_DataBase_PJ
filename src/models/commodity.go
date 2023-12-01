package models

type Commodity struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	DefaultName    string `json:"default_name" gorm:"not null;size:64"`
	ProduceAt      MyTime `json:"produce_at" gorm:"not null"`
	ProduceAddress string `json:"produce_address" gorm:"not null;size:64"`
	Category       string `json:"category" gorm:"not null;size:64"`
}
