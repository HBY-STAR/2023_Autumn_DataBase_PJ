package models

type Seller struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"user_id" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email"`
	//Age      int
	Address string `json:"address"`
}
