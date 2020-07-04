package handler

import (
	"DoctorStrange/engine"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func EngineOpen(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {

	symbol := ""
	price := 1.0

	engine.DefaultManager().AddEngine(symbol, float32(price))

	fmt.Fprint(w, "engine open")
}
