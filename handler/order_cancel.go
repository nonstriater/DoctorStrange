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

func OrderCancel(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	fmt.Fprint(w, "order cancel\n")

	symbol := ""
	side := enum.OrderSideBuy
	order := model.Order{
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action: 	enum.OrderActionCreate,
		Type:       enum.OrderTypeLimit,
		Side:      	side,
		Symbol:    	symbol,
	}

	e, _ := engine.DefaultManager().Engine(symbol)
	e.CancelOrder(order)
}