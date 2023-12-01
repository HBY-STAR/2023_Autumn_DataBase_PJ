package json

import (
	"encoding/json"
	"fmt"
	"src/models"
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
		fmt.Println(v)
		return nil
	}
	return nil
}
