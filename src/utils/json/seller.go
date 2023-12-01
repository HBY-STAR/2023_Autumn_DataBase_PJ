package json

import (
	"encoding/json"
	"fmt"
	"src/models"
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
		fmt.Println(v)
		return nil
	}
	return nil
}
