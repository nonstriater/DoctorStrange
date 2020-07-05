package model

import (
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

type QueueDirection uint8

const (
	QueueDirectionSell QueueDirection = 1 //价格升序
	QueueDirectionBuy QueueDirection = 2  //价格降序
)   

//订单队列，需要维护好订单的价格，时间顺序
//插入删除操作之后，要维护好顺序
type OrderQueue struct {
	Direction 		QueueDirection //价格排序方向
	Orders 			*sll.List 	//排序挂单列表，二维链表
	OrdersMap 		map[string][]Order//加一个map结构，提高相同价格订单查询速度
}

func NewQueue(direction QueueDirection) *OrderQueue {

	return &OrderQueue{
		Direction:  direction,
		Orders: 	sll.New(),
		OrdersMap: make(map[string][]Order),
	}
	return nil
}

func (oq *OrderQueue)AddOrder(order Order)  {
	price := order.Price
	index,exist := oq.getPriceIndex(price)
	if exist == false{
		li := sll.New()
		li.Add(order)
		oq.Orders.Insert(index,li)
		return
	}

	li,_ := oq.Orders.Get(index)
	lii,_ := li.(*sll.List)
	lii.Add(order)

	return
}

func (oq *OrderQueue)RemoveOrder(order Order)  {
	if oq.Orders.Size() == 0 {
		return
	}

	//循环遍历
	for i,li := range oq.Orders.Values() {
		lii,_ := li.(*sll.List)
		for j, _ := range lii.Values() {
			o := orderAt(oq.Orders, i, j)
			if order.ID == o.ID {
				lii.Remove(j)
				if lii.Size() == 0 {//删除后如果空了
					oq.Orders.Remove(i)
				}
				break;
			}
		}
	}
}

func (oq *OrderQueue)GetHeaderOrder() *Order {
	if oq.Orders.Size() == 0 {
		return nil
	}

	order := orderAt(oq.Orders, 0, 0)
	return &order
}

func (oq *OrderQueue)PopHeaderOrder()  *Order{
	if oq.Orders.Size() == 0 {
		return nil
	}

	li,_ := oq.Orders.Get(0)
	lii,_ := li.(*sll.List)
	lii.Remove(0)
	if lii.Size() == 0 {//删除后如果空了
		oq.Orders.Remove(0)
	}

	return nil
}

//返回ture 代表价格 已经存在，返回对应的index
//返回false 代表价格 不存在，返回 index 小
func (oq *OrderQueue)getPriceIndex(price float32)  (int, bool){
	list := oq.Orders.Values()
	if len(list) == 0 {
		return 0, false
	}

	switch oq.Direction {
	case QueueDirectionBuy:
		return getPriceIndexAtBuyQueue(oq.Orders, price)
	case QueueDirectionSell:
		return getPriceIndexAtSellQueue(oq.Orders, price)
	}

	return 0,false
}

func getPriceIndexAtSellQueue(orders *sll.List, price float32) (int, bool){
	list := orders.Values()
	if len(list) == 0 {
		return 0, false
	}

	low := orderAt(orders, 0,0)
	if price < low.Price {
		return 0, false
	}


	high := orderAt(orders, orders.Size() -1, 0)
	if price > high.Price {
		return orders.Size(), false
	}

	//二分查找算法
	for i,_ := range list {
		order := orderAt(orders, i,0)

		if order.Price == price {
			return i, true
		}

		if order.Price > price {
			return i, false
		}
	}

	return 0,false
}

func getPriceIndexAtBuyQueue(orders *sll.List, price float32) (int, bool){
	list := orders.Values()
	if len(list) == 0 {
		return 0, false
	}

	high := orderAt(orders, 0, 0)
	if price > high.Price {
		return 0, false
	}

	low := orderAt(orders, orders.Size()-1,0)
	if price < low.Price {
		return orders.Size(), false
	}

	//二分查找算法
	for i,_ := range list {
		order := orderAt(orders, i,0)

		if order.Price == price {
			return i, true
		}

		if order.Price < price {
			return i, false
		}
	}

	return 0,false
}

func orderAt(orders *sll.List, i int, j int) Order{
	li,_ := orders.Get(i)
	lii,_ := li.(*sll.List)

	obj,_ := lii.Get(j)
	order,_ := obj.(Order)
	return order
}


