package model

type PriceDirectionSide uint8

const (
	PriceDirectionASC PriceDirectionSide = 1
	PriceDirectionDESC PriceDirectionSide = 2
)   

//订单队列，需要维护好订单的价格，时间顺序
type OrderQueue struct {
	Direction PriceDirectionSide //价格排序方向
	Orders *[][]Order //排序挂单列表，二维链表
	elementMap map[string][]Order//加一个map结构，提高相同价格订单查询速度
}

func New(direction uint8)  {

}

func AddOrder(order Order)  {

}

func RemoveOrder(order Order)  {
	
}

func GetHeaderOrder()  {
	
}

func PopHeaderOrder()  {
	
}