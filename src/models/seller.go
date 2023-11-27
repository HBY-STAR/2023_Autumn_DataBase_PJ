package models

import "gorm.io/gorm"

type Seller struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"user_id" gorm:"unique;not null;size:64"`
	Password string `json:"password" gorm:"not null;size:64"`
	Email    string `json:"email" gorm:"size:64"`
	//Age      int
	Address string `json:"address" gorm:"not null;size:64"`
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
