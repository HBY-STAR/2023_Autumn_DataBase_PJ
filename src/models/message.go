package models

import "time"

// user_id		commodity_id	seller_id		platform_id		current_price
type Message struct {
	ID              int `json:"id" gorm:"primaryKey"`
	User            *User
	UserID          int `json:"user_id"`
	CommodityItem   *CommodityItem
	CommodityItemID int       `json:"commodity_item_id"`
	CurrentPrice    float64   `json:"current_price"`
	CreateAt        time.Time `json:"create_at"`
	//PriceLimit   float64
}
