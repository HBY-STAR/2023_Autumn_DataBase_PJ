package models

type Favorite struct {
	User            *User          `gorm:"ForeignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID          int            `json:"user_id" gorm:"PrimaryKey; not null"`
	CommodityItem   *CommodityItem `gorm:"ForeignKey:CommodityItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommodityItemID int            `json:"commodity_item_id" gorm:"PrimaryKey; not null"`
	PriceLimit      float32        `json:"price_limit" gorm:"not null;default:0"`
	UpdateAt        MyTime         `json:"update_at"`
}
