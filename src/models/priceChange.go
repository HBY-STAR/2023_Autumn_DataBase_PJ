package models

import "time"

type PriceChange struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	CommodityItem   *CommodityItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommodityItemID int            `json:"commodity_item_id" gorm:"not null"`
	NewPrice        float64        `json:"new_price" gorm:"not null"`
	UpdateAt        time.Time      `json:"update_at"`
}
