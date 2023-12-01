package json

import (
	"encoding/json"
	"fmt"
	"src/models"
	"src/utils"
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
		v.Password = utils.MakePassword(v.Password)
		//fmt.Println(v)
		//return nil
		err = models.DB.Create(&v).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
