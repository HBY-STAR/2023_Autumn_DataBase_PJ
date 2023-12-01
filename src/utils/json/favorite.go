package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parseFavorite() error {
	data, err := ReadString(path + "favorite.json")
	if err != nil {
		return err
	}
	var favorite []models.Favorite
	err = json.Unmarshal([]byte(data), &favorite)
	if err != nil {
		return err
	}
	for _, v := range favorite {
		fmt.Println(v)
		return nil
	}
	return nil
}
