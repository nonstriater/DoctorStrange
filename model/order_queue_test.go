package model

import (
	"DoctorStrange/enum"
	"fmt"
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
	"testing"
	"time"
)

func TestOrderQueue_AddOrder_buy(t *testing.T) {

	//买盘
	oq := NewQueue(QueueDirectionBuy)
	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideBuy,
		Symbol:    "btc/usdt",
		Price:     100,
		Amount:    1,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideBuy,
		Symbol:    "btc/usdt",
		Price:     105,
		Amount:    2,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideBuy,
		Symbol:    "btc/usdt",
		Price:     104,
		Amount:    3,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideBuy,
		Symbol:    "btc/usdt",
		Price:     101,
		Amount:    3,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideBuy,
		Symbol:    "btc/usdt",
		Price:     104,
		Amount:    4,
	})

	fmt.Printf("%s\n", oq.Orders.String())

	order := oq.PopHeaderOrder()
	fmt.Printf("%#v\n", order)
	fmt.Printf("%s\n", oq.Orders.String())

	order = oq.PopHeaderOrder()
	fmt.Printf("%#v\n", order)
	fmt.Printf("%s\n", oq.Orders.String())
}

func TestOrderQueue_AddOrder_sell(t *testing.T) {

	//买盘
	oq := NewQueue(QueueDirectionSell)
	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideSell,
		Symbol:    "btc/usdt",
		Price:     100,
		Amount:    1,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideSell,
		Symbol:    "btc/usdt",
		Price:     105,
		Amount:    2,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideSell,
		Symbol:    "btc/usdt",
		Price:     104,
		Amount:    3,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideSell,
		Symbol:    "btc/usdt",
		Price:     101,
		Amount:    3,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideSell,
		Symbol:    "btc/usdt",
		Price:     104,
		Amount:    4,
	})

	fmt.Printf("%s\n", oq.Orders.String())

	order := oq.PopHeaderOrder()
	fmt.Printf("%#v\n", order)
	fmt.Printf("%s\n", oq.Orders.String())

	order = oq.PopHeaderOrder()
	fmt.Printf("%#v\n", order)
	fmt.Printf("%s\n", oq.Orders.String())
}

func Test_orderAt(t *testing.T){
	//买盘
	oq := NewQueue(QueueDirectionBuy)
	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideBuy,
		Symbol:    "btc/usdt",
		Price:     100,
		Amount:    1,
	})

	oq.AddOrder(&Order{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSideBuy,
		Symbol:    "btc/usdt",
		Price:     102,
		Amount:    2,
	})

	orders := oq.Orders

	li,_ := orders.Get(0)
	lii,_ := li.(*sll.List)

	obj,_ := lii.Get(0)
	order,_ := obj.(*Order)

	order.Amount = 50
}

