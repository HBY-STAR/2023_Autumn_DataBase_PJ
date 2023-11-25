package favorite

type AddFavoriteModel struct {
	ItemId int `json:"item_id" validate:"required"`
}

type PriceLimitModel struct {
	PriceLimit int `json:"price_limit"`
	ItemId     int `json:"item_id"`
}
