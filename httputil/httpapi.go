// Copyright 2020 The Matrix.org Foundation C.I.C.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httputil

import (
	"kens/demo/auth"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"kens/demo/util/enty_logger"
	"net/http"
	"time"
)

// BasicAuth is used for authorization on /metrics handlers
type BasicAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

/**
登录验证
MakeUserAuthAPI turns a util.JSONRequestHandler function into an http.Handler which authenticates the request.
*/
func MakeUserAuthAPI(
	metricsName string,
	db *storage.Database,
	f func(*http.Request, *types.User) util.JSONResponse,
) http.Handler {
	h := func(req *http.Request) util.JSONResponse {
		start := time.Now()
		account, jsonRes := auth.VerifyUserFromRequest(req, db)
		if jsonRes != nil {
			return *jsonRes
		}
		// add the user ID to the logger
		logger := util.GetLogger((req.Context()))
		//logger = logger.WithField("user_id", device.UserID)
		req = req.WithContext(util.ContextWithLogger(req.Context(), logger))
		ret := f(req, account)
		enty_logger.Debug(req.RequestURI, " cost:", time.Since(start))
		return ret
	}
	return MakeExternalAPI(metricsName, h)
}

/**
登录验证
MakeClientAuthAPI turns a util.JSONRequestHandler function into an http.Handler which authenticates the request.
*/
func MakeClientAuthAPI(
	metricsName string,
	db *storage.Database,
	f func(*http.Request, *types.Client) util.JSONResponse,
) http.Handler {
	h := func(req *http.Request) util.JSONResponse {
		start := time.Now()
		account, jsonRes := auth.VerifyClientFromRequest(req, db)
		if jsonRes != nil {
			return *jsonRes
		}
		// add the user ID to the logger
		logger := util.GetLogger((req.Context()))
		//logger = logger.WithField("user_id", device.UserID)
		req = req.WithContext(util.ContextWithLogger(req.Context(), logger))
		ret := f(req, account)
		enty_logger.Debug(req.RequestURI, " cost:", time.Since(start))
		return ret
	}
	return MakeExternalAPI(metricsName, h)
}

/**
小程序用户登录验证
MakeCustomerAuthAPI turns a util.JSONRequestHandler function into an http.Handler which authenticates the request.
*/
func MakeCustomerAuthAPI(
	metricsName string,
	db *storage.Database,
	f func(*http.Request, *types.Customer) util.JSONResponse,
) http.Handler {
	h := func(req *http.Request) util.JSONResponse {
		start := time.Now()
		account, jsonRes := auth.VerifyCustomerFromRequest(req, db)
		if jsonRes != nil {
			return *jsonRes
		}
		// add the user ID to the logger
		logger := util.GetLogger((req.Context()))
		//logger = logger.WithField("user_id", device.UserID)
		req = req.WithContext(util.ContextWithLogger(req.Context(), logger))
		ret := f(req, account)
		enty_logger.Debug(req.RequestURI, " cost:", time.Since(start))
		return ret
	}
	return MakeExternalAPI(metricsName, h)
}

/**
免登录验证
MakeExternalAPI turns a util.JSONRequestHandler function into an http.Handler.
This is used for APIs that are called from the internet.
*/
func MakeExternalAPI(metricsName string, f func(*http.Request) util.JSONResponse) http.Handler {

	h := util.MakeJSONAPI(util.NewJSONRequestHandler(f))
	corfunc := func(w http.ResponseWriter, req *http.Request) {
		enty_logger.Debug("request path:", req.RequestURI)
		nextWriter := w
		h.ServeHTTP(nextWriter, req)
	}
	httpHandler := http.HandlerFunc(corfunc)

	return httpHandler
}
