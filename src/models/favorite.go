package models

import (
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
)

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

func (favorite *Favorite) Create() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(favorite).Error
	})
}

func (favorite *Favorite) Update() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Updates(&favorite)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return common.NotFound("Favorite not found")
		}
		return nil
	})
}

//func (favorite *Favorite) Save() error {
//	return DB.Transaction(func(tx *gorm.DB) error {
//		return tx.Save(favorite).Error
//	})
//}

// Delete check the user id and item id
func (favorite *Favorite) Delete() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("user_id = ? AND commodity_item_id = ?", favorite.UserID, favorite.CommodityItemID).Delete(&favorite)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return common.NotFound("Favorite not found")
		}
		return nil
	})
}
