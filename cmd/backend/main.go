package main

import (
	"github.com/gorilla/mux"
	"kens/demo/routing"
	"kens/demo/schedule"
	"kens/demo/storage"
	"kens/demo/util"
	"kens/demo/util/enty_logger"
	"kens/demo/util/environment"
	"math/rand"
	"net/http"
)

func main() {
	environment.Init()
	enty_logger.Init()
	rand.Seed(util.TimeNow().UTC().UnixNano())

	//var cli *payment.PayClient
	routers := mux.NewRouter()
	db, err := storage.NewDatabase()
	if err != nil {
		enty_logger.Info("err:", err)
		panic("db failed init")
	}
	// run schedule
	schedule.Run(db)
	routing.Setup(routers, db)
	err = http.ListenAndServe("0.0.0.0:9201", routers)
	if err != nil {
		panic("error" + err.Error())
	}
}
