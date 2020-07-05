package handler

import (
	"DoctorStrange/enum"
	"DoctorStrange/errorcode"
	"DoctorStrange/model"
	"DoctorStrange/process"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func OrderCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {

	w.Header().Set("Content-Type", "application/json")

	v := r.URL.Query()
	symbol := v.Get("symbol")
	if len(symbol) == 0 {
		w.Write(errorcode.ErrorCodeParamInvalidSymbol.ToJson())
		return
	}

	side := v.Get("side")
	price := v.Get("price")
	amount := v.Get("amount")

	s, _  := strconv.ParseInt(side, 10, 32 )
	p, _ := strconv.ParseFloat(price, 32)
	a, _ := strconv.ParseFloat(amount, 32)

	order := model.Order{
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action:    enum.OrderActionCreate,
		Type:      enum.OrderTypeLimit,
		Side:      enum.OrderSide(s),
		Symbol:    symbol,
		Price:     float32(p),
		Amount:    float32(a),
	}

	//sava order

	process.Dispatch(order)

	fmt.Fprint(w, "order add\n")
}
