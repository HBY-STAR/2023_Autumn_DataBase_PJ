package models

import "gorm.io/gorm"

type Commodity struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	DefaultName    string `json:"default_name" gorm:"not null;size:64"`
	ProduceAt      MyTime `json:"produce_at" gorm:"not null"`
	ProduceAddress string `json:"produce_address" gorm:"not null;size:64"`
	Category       string `json:"category" gorm:"not null;size:64"`
}

func GetCommodities() ([]Commodity, error) {
	var commodities []Commodity
	err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.Find(&commodities).Error
	})
	if err != nil {
		return nil, err
	}
	return commodities, nil
}

func GetCommodityByID(commodityID int) (commodity *Commodity, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Take(&commodity, commodityID).Error
	})
	return
}

func (commodity *Commodity) Create() (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&commodity).Error
	})
	return
}

func DeleteCommodityByID(commodityID int) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&Commodity{}, commodityID).Error
	})
	return
}

func (commodity *Commodity) Update() (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Updates(&commodity).Error
	})
	return
}
