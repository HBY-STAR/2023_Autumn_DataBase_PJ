package models

type Commodity struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	DefaultName string `json:"default_name"`
	Type        string `json:"type"`
}
