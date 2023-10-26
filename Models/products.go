package models

import "time"

type Products struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Created_At   time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serialnumber"`
}
