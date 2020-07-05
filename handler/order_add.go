package handler

import (
	"DoctorStrange/enum"
	"DoctorStrange/errorcode"
	"DoctorStrange/model"
	"DoctorStrange/process"
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
	if len(side) == 0 {
		w.Write(errorcode.ErrorCodeParamInvalidSide.ToJson())
		return
	}

	price := v.Get("price")
	if len(price) == 0 {
		w.Write(errorcode.ErrorCodeParamInvalidPrice.ToJson())
		return
	}

	amount := v.Get("amount")
	if len(amount) == 0 {
		w.Write(errorcode.ErrorCodeParamInvalidAmount.ToJson())
		return
	}

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

	w.Write(errorcode.OK.ToJson())
}
