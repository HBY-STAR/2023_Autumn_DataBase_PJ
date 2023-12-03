package models

import (
	"errors"
	"gorm.io/gorm"
	"sort"
	"time"
)

type PriceChange struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	CommodityItem   *CommodityItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommodityItemID int            `json:"commodity_item_id" gorm:"not null"`
	NewPrice        float32        `json:"new_price" gorm:"not null"`
	UpdateAt        MyTime         `json:"update_at"`
}

type ByUpdateAt []PriceChange

func (a ByUpdateAt) Len() int           { return len(a) }
func (a ByUpdateAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpdateAt) Less(i, j int) bool { return a[i].UpdateAt.Time.Before(a[j].UpdateAt.Time) }

func GetPriceChangeById(commodityItemID int, timeStart time.Time, timeEnd time.Time) (priceChanges []PriceChange, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.
			Where("commodity_item_id = ? AND update_at BETWEEN ? AND ?", commodityItemID, timeStart, timeEnd).
			//Preload("CommodityItem").
			Find(&priceChanges).Error
	})
	sort.Sort(ByUpdateAt(priceChanges))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return priceChanges, nil
	}
	return
}

func (priceChange *PriceChange) Update() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Updates(&priceChange)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}

func IsPriceChangeToday(commodityItemID int) (bool, error) {
	var priceChange PriceChange
	err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("commodity_item_id = ? AND DATE(update_at) = DATE(NOW())", commodityItemID).First(&priceChange).Error
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreatePriceChanges(priceChanges []PriceChange) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&priceChanges).Error
	})
}
