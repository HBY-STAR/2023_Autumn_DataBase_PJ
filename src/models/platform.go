package models

type Platform struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Country string `json:"country"`
}
