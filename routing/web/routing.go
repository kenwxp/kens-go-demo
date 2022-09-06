package web

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
	subRouter := router.PathPrefix("/web").Subrouter()

	//登录外
	subRouter.Handle("/login",
		httputil.MakeExternalAPI("Login", func(req *http.Request) util.JSONResponse {
			return Login(req, db)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	subRouter.Handle("/password/change",
		httputil.MakeUserAuthAPI("PasswordChange", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return ChangePassword(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//登录内查询客户列表
	subRouter.Handle("/client/list",
		httputil.MakeUserAuthAPI("ClientList", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return GetClientListWithCondition(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//登录内新增客户
	subRouter.Handle("/client/add",
		httputil.MakeUserAuthAPI("AddClient", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return AddClient(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//登录内删除客户
	subRouter.Handle("/client/delete",
		httputil.MakeUserAuthAPI("DeleteClient", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return DeleteClient(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//登录内编辑客户
	subRouter.Handle("/client/edit",
		httputil.MakeUserAuthAPI("EditClient", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return EditClient(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//登录内重置客户密码
	subRouter.Handle("/client/reset",
		httputil.MakeUserAuthAPI("ResetClientPassword", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return ResetClientPassword(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//登录内查询角色列表
	subRouter.Handle("/role/list",
		httputil.MakeUserAuthAPI("RoleList", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return GetRoleList(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//登录内查询用户列表
	subRouter.Handle("/user/list",
		httputil.MakeUserAuthAPI("UserList", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return GetUserList(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//登录内新增角色
	subRouter.Handle("/user/add",
		httputil.MakeUserAuthAPI("AddUser", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return AddUser(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//修改用户接口
	subRouter.Handle("/user/edit",
		httputil.MakeUserAuthAPI("EditUser", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return EditUser(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//登录内删除角色
	subRouter.Handle("/user/delete",
		httputil.MakeUserAuthAPI("DeleteUser", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return DeleteUser(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//登录内重置用户密码
	subRouter.Handle("/user/reset",
		httputil.MakeUserAuthAPI("ResetUserPassword", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return ResetUserPassword(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	//登录内新增角色
	subRouter.Handle("/role/add",
		httputil.MakeUserAuthAPI("AddRole", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return AddRole(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//登录内删除角色
	subRouter.Handle("/role/delete",
		httputil.MakeUserAuthAPI("DeleteRole", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return DeleteRole(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//登录内编辑角色
	subRouter.Handle("/role/edit",
		httputil.MakeUserAuthAPI("EditRole", db, func(req *http.Request, user *types.User) util.JSONResponse {
			return EditRole(req, db, user)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
}
