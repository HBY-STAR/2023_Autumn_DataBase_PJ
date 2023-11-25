package models

import "gorm.io/gorm"

type Seller struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"user_id" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email"`
	//Age      int
	Address string `json:"address"`
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
