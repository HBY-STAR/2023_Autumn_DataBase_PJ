package json

import (
	"encoding/json"
	"fmt"
	"src/models"
	"src/utils"
)

func parseSeller() error {
	data, err := ReadString(path + "seller.json")
	if err != nil {
		return err
	}
	var sellers []models.Seller
	err = json.Unmarshal([]byte(data), &sellers)
	if err != nil {
		return err
	}
	for _, v := range sellers {
		v.Password = utils.MakePassword(v.Password)
	}
	err = models.DB.Create(&sellers).Error
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
