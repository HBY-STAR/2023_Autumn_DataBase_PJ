package json

import "src/models"

var path = "../../../GenerateData/Data/"

// var users []map[string]interface{}
type ByUpdateAt []models.PriceChange

func (a ByUpdateAt) Len() int           { return len(a) }
func (a ByUpdateAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpdateAt) Less(i, j int) bool { return a[i].UpdateAt.Time.Before(a[j].UpdateAt.Time) }

type ByCreatedAt []models.Message

func (a ByCreatedAt) Len() int           { return len(a) }
func (a ByCreatedAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreatedAt) Less(i, j int) bool { return a[i].CreateAt.Time.Before(a[j].CreateAt.Time) }
