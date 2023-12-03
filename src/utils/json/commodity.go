package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parseCommodity() error {
	data, err := ReadString(path + "commodity.json")
	if err != nil {
		return err
	}
	var commodities []models.Commodity
	err = json.Unmarshal([]byte(data), &commodities)
	if err != nil {
		return err
	}

	err = models.DB.Create(&commodities).Error
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
