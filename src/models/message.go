package models

import "time"

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
