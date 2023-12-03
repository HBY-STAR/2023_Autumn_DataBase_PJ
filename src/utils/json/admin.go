package json

import (
	"encoding/json"
	"fmt"
	"src/models"
	"src/utils"
)

func parseAdmin() error {
	data, err := ReadString(path + "admin.json")
	if err != nil {
		return err
	}
	var admins []models.Admin
	err = json.Unmarshal([]byte(data), &admins)
	if err != nil {
		return err
	}
	for _, v := range admins {
		v.Password = utils.MakePassword(v.Password)
		err = models.DB.Create(&v).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}