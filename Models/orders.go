package models

import "time"

type Orders struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Created_At   time.Time
	ProductRefer int      `json:"product_id"`
	Product      Products `gorm:"foreignKey:ProductRefer"`
	UserRefer    int      `json:"user_id"`
	User         User     `gorm:"foreignKey:UserRefer"`
}
