package models

import "time"

type PriceChange struct {
	ID              int `json:"id" gorm:"primaryKey"`
	CommodityItem   *CommodityItem
	CommodityItemID int       `json:"commodity_item_id"`
	NewPrice        float64   `json:"new_price"`
	UpdateAt        time.Time `json:"update_at"`
}
