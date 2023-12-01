package json

import (
	"encoding/json"
	"fmt"
	"sort"
	"src/models"
)

func parseMessage() error {
	data, err := ReadString(path + "message.json")
	if err != nil {
		return err
	}
	var message []models.Message
	err = json.Unmarshal([]byte(data), &message)
	if err != nil {
		return err
	}
	sort.Sort(ByCreatedAt(message))
	for _, v := range message {
		err = models.DB.Create(&v).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
