package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null;size:64"`
	Password string `json:"password;not null;size:64"`
	Email    string `json:"email" gorm:"size:64"`
	Age      int8   `json:"age" gorm:"not null"`
	Gender   bool   `json:"gender" gorm:"not null"`
	Phone    string `json:"phone;not null" gorm:"type:char(13)"`
}

func (user *User) LoadUserByID(userID int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Take(&user, userID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// insert user if not found
				user.ID = userID
				err = tx.Create(&user).Error
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		return nil
	})
}
