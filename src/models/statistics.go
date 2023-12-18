package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func GetFavoriteStatisticsAll() (favoriteStatisticsResponse []FavoriteStatisticsResponse, err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		var favStats []FavoriteStatistics
		err = tx.Model(&Favorite{}).
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
			err = tx.
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
	})
	return
}

func GetFavoriteStatisticsSome(gender *bool, ageStart int, ageEnd int) (favoriteStatisticsResponse []FavoriteStatisticsResponse, err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		var favStats []FavoriteStatistics
		if gender != nil {
			err = tx.Model(&Favorite{}).
				Joins("JOIN user ON user.id = favorite.user_id").
				Where("user.gender = ? AND user.age BETWEEN ? AND ?", gender, ageStart, ageEnd).
				Limit(10).
				Select("count(*) as count, commodity_item_id").
				Group("commodity_item_id").
				Order("count DESC").
				Scan(&favStats).Error
		} else {
			err = tx.Model(&Favorite{}).
				Joins("JOIN user ON user.id = favorite.user_id").
				Where("user.age BETWEEN ? AND ?", ageStart, ageEnd).
				Limit(10).
				Select("count(*) as count, commodity_item_id").
				Group("commodity_item_id").
				Order("count DESC").
				Scan(&favStats).Error
		}
		if err != nil {
			return
		}
		for _, stat := range favStats {
			var commodityItem CommodityItem
			err = tx.
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
	})
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

func GetPriceStatisticsHistory(commodityID int, timeStart time.Time, timeEnd time.Time) (priceStatisticsResponse []PriceStatisticsResponse, err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		var items []CommodityItem
		err = tx.Model(&CommodityItem{}).
			Where("commodity_id = ?", commodityID).
			Preload(clause.Associations).
			Find(&items).Error
		if err != nil {
			return
		}
		for _, item := range items {
			var PriceVariance float32
			err = tx.Model(&PriceChange{}).
				Where("commodity_item_id = ? AND update_at BETWEEN ? AND ?", item.ID, timeStart, timeEnd).
				Group("commodity_item_id").
				Select("MAX(new_price) - MIN(new_price) as price_variance").
				Find(&PriceVariance).Error
			// SELECT MAX(new_price) - MIN(new_price) AS price_variance FROM price_change WHERE commodity_item_id = 1 GROUP BY commodity_item_id
			if err != nil {
				return
			}
			priceStatisticsResponse = append(priceStatisticsResponse, PriceStatisticsResponse{
				CommodityItem: item,
				PriceVariance: PriceVariance,
			})
		}
		return
	})
	// if transfer into a sql statement, it will be:
	// SELECT commodity_item_id, MAX(new_price) - MIN(new_price) AS price_variance FROM price_change WHERE commodity_item_id IN (SELECT id FROM commodity_item WHERE commodity_id = 1) AND update_at BETWEEN '2021-05-01 00:00:00' AND '2025-05-31 00:00:00' GROUP BY commodity_item_id
	// select MAX(new_price) - MIN(new_price) as price_variance, commodity_item_id from commodity_item join price_change on commodity_item.id = price_change.commodity_item_id where commodity_id = 54 AND price_change.update_at BETWEEN "2022-01-01 09:26:50.000" AND "2024-01-01 09:26:50.000" Group By commodity_item_id
	return
}
