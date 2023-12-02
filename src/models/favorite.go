package models

import "gorm.io/gorm"

type Favorite struct {
	User            *User          `gorm:"ForeignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID          int            `json:"user_id" gorm:"PrimaryKey; not null;autoIncrement:false"`
	CommodityItem   *CommodityItem `gorm:"ForeignKey:CommodityItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommodityItemID int            `json:"commodity_item_id" gorm:"PrimaryKey; not null;autoIncrement:false"`
	PriceLimit      float32        `json:"price_limit" gorm:"not null;default:0"`
	UpdateAt        MyTime         `json:"update_at"`
}

func GetFavoritesByUserID(userID int) (favorites []Favorite, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Preload("CommodityItem").Where("user_id = ?", userID).Find(&favorites).Error
	})
	return favorites, err
}
