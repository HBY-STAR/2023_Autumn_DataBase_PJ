package models

import "time"

type Favorite struct {
	User            *User          `gorm:"ForeignKey:UserID"`
	UserID          int            `json:"user_id" gorm:"PrimaryKey"`
	CommodityItem   *CommodityItem `gorm:"ForeignKey:CommodityItemID"`
	CommodityItemID int            `json:"commodity_item_id" gorm:"PrimaryKey"`
	PriceLimit      float64        `json:"price_limit"`
	CreateAt        time.Time      `json:"create_at"`
}
