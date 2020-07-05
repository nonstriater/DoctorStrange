package handler

import (
	"DoctorStrange/engine"
	"DoctorStrange/enum"
	"DoctorStrange/model"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func OrderCancel(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {

	v := r.URL.Query()
	oid := v.Get("orderId")
	orderId,_ := strconv.ParseUint(oid,10,64)

	symbol := v.Get("symbol")
	side := v.Get("side")
	s, _  := strconv.ParseInt(side, 10, 32 )
	order := model.Order{
		ID:			uint64(orderId),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Action: 	enum.OrderActionCancel,
		Type:       enum.OrderTypeLimit,
		Side:      	enum.OrderSide(s),
		Symbol:    	symbol,
	}


	fmt.Fprint(w, "order cancel\n")
}
