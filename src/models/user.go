package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
	Email    string `json:"email"`
	Age      int8   `json:"age"`
	Gender   bool   `json:"gender"`
}
