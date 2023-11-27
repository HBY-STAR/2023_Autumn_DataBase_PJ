package models

import "time"

type Message struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	User            *User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID          int            `json:"user_id" gorm:"not null"`
	CommodityItem   *CommodityItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommodityItemID int            `json:"commodity_item_id"` //如果不constraint join时要注意
	CurrentPrice    float64        `json:"current_price" gorm:"not null"`
	CreateAt        time.Time      `json:"create_at"`
	//PriceLimit   float64
}
