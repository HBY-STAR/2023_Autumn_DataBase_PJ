package favorite

type AddFavoriteModel struct {
	ItemId int `json:"item_id" validate:"required"`
}

type PriceLimitModel struct {
	PriceLimit float32 `json:"price_limit"`
	ItemId     int     `json:"item_id"`
}
