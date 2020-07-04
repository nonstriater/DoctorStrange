package model

import (
	"DoctorStrange/enum"
)


//委托单账本
type OrderBook struct {

	BuyOrderQueue   *OrderQueue  //买单队列
	SellOrderQueue  *OrderQueue  //卖单队列
}

func (ob *OrderBook)AddOrder (o Order)  {
	switch o.Side {
	case enum.OrderSideSell:
		ob.SellOrderQueue.AddOrder(o)
	case enum.OrderSideBuy:
		ob.BuyOrderQueue.AddOrder(o)
	}
}

func (ob *OrderBook)CancelOrder (o Order)  {
	switch o.Side {
	case enum.OrderSideSell:
		ob.SellOrderQueue.RemoveOrder(o)
	case enum.OrderSideBuy:
		ob.BuyOrderQueue.RemoveOrder(o)
	}
}

func (ob *OrderBook)GetHeaderBuyOrder()  *Order{
	return ob.BuyOrderQueue.GetHeaderOrder()
}

func (ob *OrderBook)GetHeaderSellOrder()  *Order{
	return ob.SellOrderQueue.GetHeaderOrder()
}

func (ob *OrderBook)RemoveHeaderBuyOrder() *Order{
	return ob.BuyOrderQueue.PopHeaderOrder()
}

func (ob *OrderBook)RemoveHeaderSellOrder() *Order{
	return ob.SellOrderQueue.PopHeaderOrder()
}

