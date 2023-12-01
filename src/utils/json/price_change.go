package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parsePriceChange() error {
	data, err := ReadString(path + "price_change.json")
	if err != nil {
		return err
	}
	var changes []models.PriceChange
	err = json.Unmarshal([]byte(data), &changes)
	if err != nil {
		return err
	}
	for _, v := range changes {
		err = models.DB.Create(&v).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
