package models

import (
	"bytes"
	"fmt"
	"time"
)

type MyTime struct {
	time.Time
}

type Commodity struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	DefaultName    string `json:"default_name" gorm:"not null;size:64"`
	ProduceAt      MyTime `json:"produce_at" gorm:"not null"`
	ProduceAddress string `json:"produce_address" gorm:"not null;size:64"`
	Category       string `json:"category" gorm:"not null;size:64"`
}

var formatTime = "2006-01-02 15:04:05"

func (t *MyTime) UnmarshalJSON(b []byte) (err error) {
	b = bytes.Trim(b, "\"")
	tt, err := time.Parse(formatTime, string(b))
	*t = MyTime{tt}
	//now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), time.Local)
	return
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format(formatTime))), nil
}
