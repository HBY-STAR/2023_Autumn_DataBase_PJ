package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommodityItem struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Commodity   *Commodity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommodityID int        `json:"commodity_id" gorm:"not null"`
	Platform    *Platform  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlatformID  int        `json:"platform_id" gorm:"not null"`
	Seller      *Seller    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SellerID    int        `json:"seller_id" gorm:"not null"`
	ItemName    string     `json:"item_name" gorm:"not null;size:64"`
	Price       float32    `json:"price" gorm:"check:price > 0; not null"`
	UpdateAt    MyTime     `json:"update_at"`
}

func (item *CommodityItem) GetItemByID(ID int) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.Take(&item, ID).Error
	})
	return err
}

func GetItems() ([]CommodityItem, error) {
	var items []CommodityItem
	err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.Preload(clause.Associations).Find(&items).Error
	})
	if err != nil {
		return nil, err
	}
	//for _, it := range items {
	//	//it.Commodity = new(Commodity)
	//	//it.Platform = new(Platform)
	//	//it.Seller = new(Seller)
	//	//_ = it.Commodity.GetCommodityByID(it.CommodityID)
	//	//_ = it.Platform.GetPlatformByID(it.PlatformID)
	//	//_ = it.Seller.GetSellerByID(it.SellerID)
	//	fmt.Println(it.CommodityID)
	//}
	return items, nil
}
