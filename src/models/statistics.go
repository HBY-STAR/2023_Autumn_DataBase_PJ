package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func GetFavoriteStatisticsAll() (favoriteStatisticsResponse []FavoriteStatisticsResponse, err error) {
	var favStats []FavoriteStatistics
	err = DB.Model(&Favorite{}).
		Limit(10).
		Select("count(*) as count, commodity_item_id").
		Group("commodity_item_id").
		Order("count desc").
		Scan(&favStats).Error
	if err != nil {
		return
	}
	for _, stat := range favStats {
		var commodityItem CommodityItem
		err = DB.
			Preload(clause.Associations).
			First(&commodityItem, stat.CommodityItemID).Error
		if err != nil {
			return
		}
		favoriteStatisticsResponse = append(favoriteStatisticsResponse, FavoriteStatisticsResponse{
			CommodityItem: &commodityItem,
			Count:         stat.Count,
		})
	}
	return
}

func GetFavoriteStatisticsSome(gender bool, ageStart int, ageEnd int) (favoriteStatisticsResponse []FavoriteStatisticsResponse, err error) {
	var favStats []FavoriteStatistics
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&Favorite{}).
			Joins("JOIN user ON user.id = favorite.user_id").
			Where("user.gender = ? AND user.age BETWEEN ? AND ?", gender, ageStart, ageEnd).
			Limit(10).
			Select("count(*) as count, commodity_item_id").
			Group("commodity_item_id").
			Order("count DESC").
			Scan(&favStats).Error
	})
	if err != nil {
		return
	}
	for _, stat := range favStats {
		var commodityItem CommodityItem
		err = DB.
			Preload(clause.Associations).
			First(&commodityItem, stat.CommodityItemID).Error
		if err != nil {
			return
		}
		favoriteStatisticsResponse = append(favoriteStatisticsResponse, FavoriteStatisticsResponse{
			CommodityItem: &commodityItem,
			Count:         stat.Count,
		})
	}
	return
	//err = DB.Transaction(func(tx *gorm.DB) error {
	//	return tx.Model(&User{}).
	//		Where("gender = ? AND age BETWEEN ? AND ?", gender, ageStart, ageEnd).
	//		Preload("Favorites").
	//		Select("count(*) as count, commodity_item_id").
	//		Group("commodity_item_id").
	//		Order("count DESC").
	//		Scan(&favStats).Error
	//})

	// Preload("CommodityItem").
	// Preload("CommodityItem.Commodity").
	// Preload("CommodityItem.Platform").
	// Preload("CommodityItem.Seller").
	//	preload will happen after the query, so it will not work if select CommodityItem
}

func GetPriceStatisticsHistory(commodityItemID int, timeStart time.Time, timeEnd time.Time) (priceStatisticsResponse []PriceStatisticsResponse, err error) {
	return
}
