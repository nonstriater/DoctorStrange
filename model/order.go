package model

import (
	"DoctorStrange/enum"
	"time"
)


type Order struct {
	ID 				uint `gorm:"primary_key" json:"id"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"update_at"`

	Action			enum.OrderAction
	Type			enum.OrderType  //订单类型： 限价单，市价单等，目前只支持限价单
	Side  			enum.OrderSide 	//订单方向  买单:1  卖单:2
	Symbol			string  //交易对,如 btd/usdt
	Price			float32 //价格
	Amount 			float32 //数量
}

