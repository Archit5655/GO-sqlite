package models

import "time"

type User struct {
	Id uint `json:"id" gorm:"primaryKey"`

	Created_At time.Time
	First_Name string `json:"firstname"`

	Last_Name string `json:"lastname"`
}
