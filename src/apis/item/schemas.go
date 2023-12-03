package item

type CreateItemModel struct {
	CommodityID int     `json:"commodity_id" validate:"required"`
	PlatformID  int     `json:"platform_id" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	ItemName    string  `json:"item_name" validate:"required"`
}

//ProduceAt      time.Time `json:"produce_at" validate:"required"`
//ProduceAddress string    `json:"produce_address" validate:"required"`

type UpdateItemModel struct {
	CommodityItemID int     `json:"commodity_item_id" validate:"required"`
	ItemName        string  `json:"item_name" validate:"required"`
	Price           float32 `json:"price" validate:"required"`
}
