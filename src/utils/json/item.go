package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parseItem() error {
	data, err := ReadString(path + "commodity_item.json")
	if err != nil {
		return err
	}
	var items []models.CommodityItem
	err = json.Unmarshal([]byte(data), &items)
	if err != nil {
		return err
	}
	for _, v := range items {
		err = models.DB.Create(&v).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
