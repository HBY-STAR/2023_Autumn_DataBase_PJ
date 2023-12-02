package models

import "gorm.io/gorm"

type Seller struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null;size:64"`
	Password string `json:"password" gorm:"not null;size:64"`
	Email    string `json:"email" gorm:"size:64"`
	//Age      int
	Address string `json:"address" gorm:"not null;size:64"`
}

func GetSellerByID(userID int) (seller *Seller, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Omit("Password").Take(&seller, userID).Error
	})
	return
}
