package process

import (
	"DoctorStrange/engine"
	"DoctorStrange/enum"
	"DoctorStrange/logger"
	"DoctorStrange/model"
)

func Dispatch(order model.Order)  {
	//对应symbol引擎是否启动
	symbol := order.Symbol
	e, _ := engine.DefaultManager().Engine(symbol)
	if e == nil {
		logger.Info("invalid engine for symbol:" + order.Symbol)
		return
	}

	switch order.Action {
	case enum.OrderActionCreate:
		e.AddOrder(order)
	case enum.OrderActionCancel:
		e.CancelOrder(order)
	}
}
