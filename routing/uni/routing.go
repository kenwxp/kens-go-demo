package uni

import (
	"github.com/gorilla/mux"
	"kens/demo/httputil"
	"kens/demo/storage"
	"kens/demo/util"
	"net/http"
)

// Setup configures the given mux with sync-server listeners
func InitRouting(
	router *mux.Router, accountDB *storage.Database,
) {
	subRouter := router.PathPrefix("/uni").Subrouter()

	subRouter.Handle("/login/check",
		httputil.MakeExternalAPI("CheckNewCustomer", func(req *http.Request) util.JSONResponse {
			return CheckNewCustomer(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	subRouter.Handle("/login",
		httputil.MakeExternalAPI("CustomerLogin", func(req *http.Request) util.JSONResponse {
			return CustomerLogin(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

}
