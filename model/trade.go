package model

import "time"

//撮合成交以后的订单成交记录
type  Trade struct {
	ID 				uint `gorm:"primary_key" json:"id"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"update_at"`

	makerId string
	takerId string
	amount float32
	price float32
}
