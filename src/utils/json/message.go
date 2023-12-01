package json

import (
	"encoding/json"
	"fmt"
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
	for _, v := range message {
		fmt.Println(v)
		return nil
	}
	return nil
}
