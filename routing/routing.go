package routing

import (
	"github.com/gorilla/mux"
	"kens/demo/httputil"
	"kens/demo/routing/demo"
	"kens/demo/storage"
)

// Setup configures the given mux with sync-server listeners

func Setup(
	router *mux.Router, db *storage.Database, wsCli *httputil.WebSocketCli,
) {
	//init miniapp routing
	//uni.InitRouting(router, accountDB)
	////init app routing
	//app.InitRouting(router, accountDB)
	////init windows routing
	//web.InitRouting(router, accountDB)
	demo.InitRouting(router, db, wsCli)
}
