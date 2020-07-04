package model

import (
	"time"
)


//委托单账本
type OrderBook struct {
	ID 				uint `gorm:"primary_key" json:"id"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"update_at"`

	BuyOrderQueue   *OrderQueue  //买单队列
	SellOrderQueue  *OrderQueue  //卖单队列
}

func (ob *OrderBook)AddOrder (o Order)  {

}

func (ob *OrderBook)CancelOrder (o Order)  {

}