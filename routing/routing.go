package routing

import (
	"github.com/gorilla/mux"
	"kens/demo/routing/app"
	"kens/demo/routing/uni"
	"kens/demo/routing/web"
	"kens/demo/storage"
)

// Setup configures the given mux with sync-server listeners

func Setup(
	router *mux.Router, accountDB *storage.Database,
) {
	//init miniapp routing
	uni.InitRouting(router, accountDB)
	//init app routing
	app.InitRouting(router, accountDB)
	//init windows routing
	web.InitRouting(router, accountDB)
}
