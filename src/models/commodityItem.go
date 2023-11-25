package models

import "time"

type CommodityItem struct {
	ID             int `json:"id" gorm:"primaryKey"`
	Commodity      *Commodity
	CommodityID    int `json:"commodity_id"`
	Platform       *Platform
	PlatformID     int `json:"platform_id"`
	Seller         *Seller
	SellerID       int       `json:"seller_id"`
	Price          float64   `json:"price"`
	UpdateAt       time.Time `json:"update_at"`
	ProduceAt      time.Time `json:"produce_at"`
	ProduceAddress string    `json:"produce_address"`
}
