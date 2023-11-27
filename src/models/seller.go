package models

import "gorm.io/gorm"

type Seller struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"user_id" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email"`
	//Age      int
	Address string `json:"address" gorm:"not null"`
}

func (user Seller) LoadSellerByID(userID int) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Take(&user, userID).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
