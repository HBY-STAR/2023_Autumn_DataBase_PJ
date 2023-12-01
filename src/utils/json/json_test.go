package json

import (
	"encoding/json"
	"fmt"
	"src/models"
	"testing"
)

func TestJson(t *testing.T) {
	_ = parseUsers()
	_ = parseSeller()
	_ = parseCommodity()
}

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
	for _, v := range commodities {
		fmt.Println(v)
		return nil
	}
	return nil
}

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

func parseUsers() error {
	data, err := ReadString(path + "user.json")
	if err != nil {
		return err
	}
	//fmt.Println("data:", data)

	//err = json.Unmarshal([]byte(data), &users)
	//fmt.Println(len(users))
	//for _, v := range users {
	//	var singleUser = models.User{
	//		ID:       0,
	//		Username: ,
	//		Password: "",
	//		Email:    "",
	//		Age:      0,
	//		Gender:   false,
	//		Phone:    "",
	//	}
	//}
	var users []models.User
	err = json.Unmarshal([]byte(data), &users)
	if err != nil {
		return err
	}
	for _, v := range users {
		fmt.Println(v)
		return nil
	}
	return nil
}
