package models

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
