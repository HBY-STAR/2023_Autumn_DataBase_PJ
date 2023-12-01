package models

type PriceChange struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	CommodityItem   *CommodityItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommodityItemID int            `json:"commodity_item_id" gorm:"not null"`
	NewPrice        float32        `json:"new_price" gorm:"not null"`
	UpdateAt        MyTime         `json:"update_at"`
}
