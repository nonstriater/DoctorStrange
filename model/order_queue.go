package model

type QueueDirection uint8

const (
	QueueDirectionSell QueueDirection = 1 //价格升序
	QueueDirectionBuy QueueDirection = 2  //价格降序
)   

//订单队列，需要维护好订单的价格，时间顺序
type OrderQueue struct {
	Direction 		QueueDirection //价格排序方向
	Orders 			*[][]Order //排序挂单列表，二维链表
	elementMap 		map[string][]Order//加一个map结构，提高相同价格订单查询速度
}

func NewQueue(direction QueueDirection) *OrderQueue {
	return nil
}

func (oq *OrderQueue)AddOrder(order Order)  {

}

func (oq *OrderQueue)RemoveOrder(order Order)  {
	
}

func (oq *OrderQueue)GetHeaderOrder() *Order {
	return nil
}

func (oq *OrderQueue)PopHeaderOrder()  *Order{
	return nil
}