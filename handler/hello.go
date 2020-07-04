package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	fmt.Fprint(w, "world\n")
	fmt.Println("world!")
}
