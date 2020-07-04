package handler

import (
	"DoctorStrange/engine"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func EngineClose(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {

	u := r.URL.Query()
	symbol := u.Get("symbol")
	engine.DefaultManager().RemoveEngine(symbol)

	fmt.Fprint(w, "engine close")
}