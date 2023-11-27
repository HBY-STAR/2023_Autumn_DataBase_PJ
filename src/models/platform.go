package models

type Platform struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"not null;size:64"`
	URL     string `json:"url" gorm:"size:64"`
	Country string `json:"country" gorm:"size:64"`
}
