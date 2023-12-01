package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

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
