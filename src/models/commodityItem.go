package models

import "time"

type CommodityItem struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Commodity   *Commodity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommodityID int        `json:"commodity_id" gorm:"not null"`
	Platform    *Platform  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlatformID  int        `json:"platform_id" gorm:"not null"`
	Seller      *Seller    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SellerID    int        `json:"seller_id" gorm:"not null"`
	ItemName    string     `json:"item_name" gorm:"not null"`
	Price       float64    `json:"price" gorm:"check:price > 0; not null"`
	UpdateAt    time.Time  `json:"update_at"`
}
