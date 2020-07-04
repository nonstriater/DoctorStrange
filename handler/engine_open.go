package handler

import (
	"DoctorStrange/engine"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func EngineOpen(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {

	u := r.URL.Query()

	symbol := u.Get("symbol")
	price := u.Get("price")//开盘价
	p,_ := strconv.ParseFloat(price, 64)

	engine.DefaultManager().AddEngine(symbol, float32(p))

	fmt.Fprint(w, "engine open")
}
