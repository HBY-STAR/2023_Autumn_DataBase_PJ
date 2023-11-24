package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	IsAdmin  bool
	Email    string
	Age      int
	Gender   bool
}
