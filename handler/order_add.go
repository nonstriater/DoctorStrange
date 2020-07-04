package handler

import (
	"DoctorStrange/engine"
	"DoctorStrange/enum"
	"DoctorStrange/model"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func OrderCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	fmt.Fprint(w, "order add\n")

	side := enum.OrderSideBuy
	symbol := ""
	price := float32(1)
	amount := float32(1)
	order := model.Order{
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action: enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      side,
		Symbol:    symbol,
		Price:     price,
		Amount:    amount,
	}

	e, _ := engine.DefaultManager().Engine(symbol)
	e.AddOrder(order)

}
