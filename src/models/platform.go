package models

import "gorm.io/gorm"

type Platform struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"not null;size:64"`
	URL     string `json:"url" gorm:"size:64"`
	Country string `json:"country" gorm:"size:64"`
}

func GetPlatforms() (platforms []Platform, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Find(&platforms).Error
	})
	if err != nil {
		return nil, err
	}
	return
}

func GetPlatformByID(platformID int) (platform *Platform, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Take(&platform, platformID).Error
	})
	if err != nil {
		return nil, err
	}
	return
}
