package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parsePlatform() error {
	data, err := ReadString(path + "platform.json")
	if err != nil {
		return err
	}
	var platforms []models.Platform
	err = json.Unmarshal([]byte(data), &platforms)
	if err != nil {
		return err
	}
	for _, v := range platforms {
		err = models.DB.Create(&v).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
