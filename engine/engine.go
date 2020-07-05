package engine

import (
	"DoctorStrange/enum"
	"DoctorStrange/logger"
	"DoctorStrange/model"
)

type engine struct {
	symbol 		string	 //交易对
	price   	float32  //最新实时成交价
	orderBook 	*model.OrderBook

	orderChan 	chan model.Order
	stop 		chan bool //引擎结束信号
}

func New(symbol string, price float32)(e *engine)  {
	if len(symbol) == 0 {
		return nil
	}

	return &engine{
		symbol:    symbol,
		price:		price,
		orderBook: &model.OrderBook{
			BuyOrderQueue:model.NewQueue(model.QueueDirectionBuy),
			SellOrderQueue:model.NewQueue(model.QueueDirectionSell),
		},
		orderChan: 	make(chan model.Order),
		stop: 		make(chan bool),
	}
}

func (e *engine) Prepare()  {

}

//price 开盘价
func (e *engine) Start()  {

	//从持久化磁盘上恢复order book
	recoverOrderBook()

	//开启撮合服务
	go e.run()

}

func (e *engine) Stop()  {
	//将当前的委托账本持久化，确保数据不丢失
	saveOrderBook()

	close(e.orderChan)
	e.stop <- true
}

func (e *engine) AddOrder(o model.Order)  {

	e.orderChan <- o
}

func (e *engine) CancelOrder(o model.Order)  {

	e.orderChan <- o
}

func (e *engine)run() {
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

	switch o.Side {
	case enum.OrderSideBuy:
		e.buyOrderMatching(o)
	case enum.OrderSideSell:
		e.sellOrderMatching(o)
	}

}

func (e *engine)dealCancelOrder(o model.Order)  {
	e.orderBook.CancelOrder(o)
}

//买单 匹配引擎策略实现
func (e *engine)buyOrderMatching(o model.Order) {

LOOP:
	oldPrice := e.price
	s := e.orderBook.GetHeaderSellOrder()
	if o.Price < s.Price {
		e.orderBook.AddOrder(o)
		return;
	}

	//new price
	newPrice := newPrice(o.Price, s.Price, oldPrice)
	e.price = newPrice

	if o.Amount == s.Amount {
		createTrade(o.ID,s.ID, oldPrice, newPrice, o.Amount )
		e.orderBook.RemoveHeaderSellOrder()
		return
	}

	if (o.Amount < s.Amount){
		createTrade(o.ID,s.ID, oldPrice, newPrice, o.Amount)
		s.Amount -= o.Amount
		return
	}

	if o.Amount > s.Amount {
		createTrade(o.ID,s.ID, oldPrice, newPrice, o.Amount)
		e.orderBook.RemoveHeaderSellOrder()
		o.Amount -= s.Amount
		goto LOOP
	}

}

//卖单 匹配引擎策略实现
func (e *engine)sellOrderMatching(o model.Order) {

LOOP:
	oldPrice := e.price
	b := e.orderBook.GetHeaderBuyOrder()
	if b.Price < o.Price {
		e.orderBook.AddOrder(o)
		return;
	}

	//new price
	newPrice := newPrice(b.Price, o.Price, oldPrice)
	e.price = newPrice

	if b.Amount == o.Amount {
		createTrade(b.ID,o.ID, oldPrice, newPrice, o.Amount )
		e.orderBook.RemoveHeaderBuyOrder()
		return
	}

	if (o.Amount < b.Amount){
		createTrade(b.ID,o.ID, oldPrice, newPrice, o.Amount)
		b.Amount -= o.Amount
		return
	}

	if o.Amount > b.Amount {
		createTrade(b.ID,o.ID, oldPrice, newPrice, b.Amount)
		e.orderBook.RemoveHeaderBuyOrder()
		o.Amount -= b.Amount
		goto LOOP
	}

}

func newPrice(buyPrice float32, sellPrice float32, oldPrice float32)  float32{
	if buyPrice < sellPrice {
		 return oldPrice
	}

	if oldPrice >= buyPrice {
		return buyPrice
	}

	if oldPrice <= sellPrice {
		 return sellPrice
	}

	if  oldPrice < buyPrice && oldPrice > sellPrice{//成交价不变
		return oldPrice
	}

	return oldPrice
}

//创建一条成交记录
func createTrade(buyOrderId uint64, sellOrderId uint64, oldPrice float32,  newPrice float32, amount float32 ) {
	logger.Info("new trade:buyOrderId=%d,sellOrderId=%d,oldPrice=%f,newPrice=%f,amount=%f",
		buyOrderId,sellOrderId,oldPrice,newPrice,amount)


}


//引擎启动时，从磁盘恢复挂单账本
func recoverOrderBook()  {

}

//引擎关闭时，持久化 order book
func saveOrderBook() {

}



