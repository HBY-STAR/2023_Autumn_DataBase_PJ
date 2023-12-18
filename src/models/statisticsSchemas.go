package models

type FavoriteStatisticsResponse struct {
	CommodityItem *CommodityItem
	Count         int `json:"count"`
}

type FavoriteStatistics struct {
	Count           int `json:"count"`
	CommodityItemID int `json:"commodity_item_id"`
}

type PriceStatisticsResponse struct {
	CommodityItem CommodityItem
	PriceVariance float32 `json:"price_variance"`
}

//type PriceStatisticsResponseCurrent struct {
//	CommodityItem *CommodityItem
//}
