package main

import (
	"DoctorStrange/engine"
	"DoctorStrange/handler"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main()  {

	initViper()
	initLog()

	initRedis()

	startRPC()

	engine.DefaultManager().Start()
}

func initViper()  {

}

func initLog()  {

}

func initRedis()  {

}

func startRPC() {

	router := httprouter.New()

	router.GET("/", handler.Hello)
	router.GET("/order/create", handler.OrderCreate)
	router.GET("/order/cancel", handler.OrderCancel)

	router.GET("/engine/open", handler.EngineOpen)
	router.GET("/engine/close", handler.EngineClose)

	log.Fatal(http.ListenAndServe(":8080", router))
}
