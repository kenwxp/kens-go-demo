package app

import (
	"github.com/gorilla/mux"
	"kens/demo/httputil"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"net/http"
)

// Setup configures the given mux with sync-server listeners
func InitRouting(
	router *mux.Router, db *storage.Database,
) {
	subRouter := router.PathPrefix("/app").Subrouter()

	//注册
	subRouter.Handle("/register",
		httputil.MakeExternalAPI("Register", func(req *http.Request) util.JSONResponse {
			return Register(req, db)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//登录外
	subRouter.Handle("/login",
		httputil.MakeExternalAPI("Login", func(req *http.Request) util.JSONResponse {
			return Login(req, db)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//修改密码
	subRouter.Handle("/password/change",
		httputil.MakeClientAuthAPI("ChangePassword", db, func(req *http.Request, client *types.Client) util.JSONResponse {
			return ChangePassword(req, db, client)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
}
