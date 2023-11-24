package models

type Seller struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string `gorm:"column:phone"`
	Email    string
	Age      int
}
