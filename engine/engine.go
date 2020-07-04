package engine

import (
	"DoctorStrange/enum"
	"DoctorStrange/model"
)

type engine struct {
	symbol 	string	 //交易对
	orderBook *model.OrderBook
	orderChan chan model.Order
	stop chan bool //结束信号
}

func New(symbol string)(e *engine)  {
	if len(symbol) == 0 {
		return nil
	}

	return &engine{
		symbol:    symbol,
		orderBook: &model.OrderBook{},
		orderChan: make(chan model.Order),
		stop : make(chan bool),
	}
}

func (e *engine) Prepare()  {

}

//price 开盘价
func (e *engine) Start(price float32)  {

	//从持久化磁盘上恢复order book
	recoverOrderBook()

	//开启撮合服务
	go e.run(price)

}

func (e *engine) Stop()  {
	//将当前的委托账本持久化，确保数据不丢失
	saveOrderBook()

	e.stop <- true
}

func (e *engine) startMetrics()  {

}

func (e *engine) AddOrder(o model.Order)  {

	e.orderChan <- o
}

func (e *engine) CancelOrder(o model.Order)  {

	e.orderChan <- o
}

func (e *engine)run(price float32) {
	for  {
		select {
		case o := <- e.orderChan:
			//等待新的买卖单加入 orderbook
			switch o.Action {
			case enum.OrderActionCreate:
				e.dealCreateOrder(o)
			case enum.OrderActionCancel:
				e.dealCancelOrder(o)
			}
		case <- e.stop:
			break
		}

	}
}

func (e *engine)dealCreateOrder(o model.Order)  {
	e.orderBook.AddOrder(o)
	e.orderMatching()
}

func (e *engine)dealCancelOrder(o model.Order)  {
	e.orderBook.CancelOrder(o)
}

//匹配引擎策略实现
func (e *engine)orderMatching() {


}

func recoverOrderBook()  {

}

func saveOrderBook() {

}



