package main

import (
	"DoctorStrange/engine"
	"DoctorStrange/handler"
	"DoctorStrange/logger"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main()  {

	initViper()
	initLog()

	initRedis()

	startMetrics()

	engine.DefaultManager().Start()

	startServer()
}

func initViper()  {

}

func initLog()  {
	logger.InitCustomLog("./logs/", "com.doctorstrange", 2)
	logger.Info("init logger")
}

func initRedis()  {

}

func startMetrics()  {

}

func startServer() {

	router := httprouter.New()

	router.GET("/", handler.Hello)
	router.GET("/order/create", handler.OrderCreate)
	router.GET("/order/cancel", handler.OrderCancel)

	router.GET("/engine/open", handler.EngineOpen)
	router.GET("/engine/close", handler.EngineClose)

	log.Fatal(http.ListenAndServe(":8080", router))
}
