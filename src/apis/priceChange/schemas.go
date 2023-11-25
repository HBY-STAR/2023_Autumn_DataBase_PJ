package priceChange

import "time"

type GetPriceChangeModel struct {
	CommodityItemID int       `json:"commodity_item_id" validate:"required"`
	TimeStart       time.Time `json:"time_start" validate:"required"`
	TimeEnd         time.Time `json:"time_end" validate:"required"`
}

type UpdatePriceChangeModel struct {
	ID              int     `json:"id"`
	CommodityItemID int     `json:"commodity_item_id"`
	NewPrice        float64 `json:"new_price"`
}
