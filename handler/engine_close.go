package handler

import (
	"DoctorStrange/engine"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func EngineClose(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {

	symbol := ""
	engine.DefaultManager().RemoveEngine(symbol)

	fmt.Fprint(w, "engine close")
}