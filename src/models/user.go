package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null;size:64"`
	Password string `json:"password" gorm:"not null;size:64"`
	Email    string `json:"email" gorm:"size:64"`
	Age      int8   `json:"age" gorm:"not null"`
	Gender   bool   `json:"gender" gorm:"not null"`
	Phone    string `json:"phone" gorm:"not null;type:char(13)"`
}

func GetUserByID(userID int) (user *User, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Omit("Password").Take(&user, userID).Error
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
	return
}
