// Copyright 2017 Vector Creations Ltd
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

// Package auth implements authentication checks and storage.
package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"kens/demo/jsonerror"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"kens/demo/util/enty_logger"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// VerifyUserFromRequest authenticates the HTTP request,
// on success returns Device of the requester.
// Finds local user or an application service user.
// Note: For an AS user, AS dummy device is returned.
// On failure returns an JSON error response which can be sent to the client.
func VerifyUserFromRequest(
	req *http.Request, db *storage.Database,
) (*types.User, *util.JSONResponse) {
	// Try to find the Application Service user
	token, err := ExtractAccessToken(req)
	enty_logger.Debug("find user token: ", token, err)
	if err != nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken(err.Error()),
		}
	}
	userPre, err := db.SelectUserByToken(req.Context(), nil, token)
	if err != nil || userPre == nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken("token is unauthorized "),
		}
	}
	enty_logger.Debug("user: ", userPre.AccNo, "-", userPre.AccName)
	mid := strings.Split(userPre.Token, ":")
	tokenPre := mid[0]
	timeStamp := mid[1]
	preTs, err := strconv.ParseInt(timeStamp, 10, 64)
	nowTs := util.TimeNow().Unix()

	// if tokenTs is exacted 1 week it's broken
	diff := nowTs - preTs
	unify := 1 * 7 * 24 * 60 * 60 * time.Second
	if time.Duration(diff*1000*1000*1000) > unify {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken("token is out of date"),
		}
	}
	nowToken := tokenPre + ":" + strconv.FormatInt(nowTs, 10)
	userPre.Token = nowToken
	err = db.UpdateUserToken(req.Context(), nil, userPre.AccNo, nowToken)
	if err != nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken(err.Error()),
		}
	}
	return userPre, nil
}

// VerifyClientFromRequest authenticates the HTTP request,
// on success returns Device of the requester.
// Finds local user or an application service user.
// Note: For an AS user, AS dummy device is returned.
// On failure returns an JSON error response which can be sent to the client.
func VerifyClientFromRequest(
	req *http.Request, db *storage.Database,
) (*types.Client, *util.JSONResponse) {
	// Try to find the Application Service user
	token, err := ExtractAccessToken(req)
	enty_logger.Debug("find client token: ", token, err)
	//enty_logger.Info("find token: ", token, err)
	if err != nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken(err.Error()),
		}
	}
	userPre, err := db.SelectClientByToken(req.Context(), nil, token)
	if err != nil || userPre == nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken("token is unauthorized"),
		}
	}
	enty_logger.Debug("client: ", userPre.ClientId, "-", userPre.ClientName)
	mid := strings.Split(userPre.Token, ":")
	tokenPre := mid[0]
	timeStamp := mid[1]
	preTs, err := strconv.ParseInt(timeStamp, 10, 64)
	nowTs := util.TimeNow().Unix()

	// if tokenTs is exacted 1 week it's broken
	diff := nowTs - preTs
	unify := 1 * 7 * 24 * 60 * 60 * time.Second
	if time.Duration(diff*1000*1000*1000) > unify {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken("token is out of date"),
		}
	}
	nowToken := tokenPre + ":" + strconv.FormatInt(nowTs, 10)
	userPre.Token = nowToken
	err = db.UpdateClientToken(req.Context(), nil, userPre.ClientId, nowToken)
	if err != nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken(err.Error()),
		}
	}
	return userPre, nil
}

// VerifyCustomerFromRequest authenticates the HTTP request,
// on success returns Device of the requester.
// Finds local user or an application service user.
// Note: For an AS user, AS dummy device is returned.
// On failure returns an JSON error response which can be sent to the client.
func VerifyCustomerFromRequest(
	req *http.Request, db *storage.Database,
) (*types.Customer, *util.JSONResponse) {
	// Try to find the Application Service user
	token, err := ExtractAccessToken(req)
	enty_logger.Debug("find customer token: ", token, err)
	if err != nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken(err.Error()),
		}
	}
	userPre, err := db.SelectCustomerByToken(req.Context(), nil, token)
	if err != nil || userPre == nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken("token is unauthorized"),
		}
	}
	enty_logger.Debug("customer: ", userPre.CustomerId, "-", userPre.Phone)
	mid := strings.Split(userPre.Token, ":")
	tokenPre := mid[0]
	timeStamp := mid[1]
	preTs, err := strconv.ParseInt(timeStamp, 10, 64)
	nowTs := util.TimeNow().Unix()

	// if tokenTs is exacted 1 week it's broken
	diff := nowTs - preTs
	unify := 1 * 7 * 24 * 60 * 60 * time.Second
	if time.Duration(diff*1000*1000*1000) > unify {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken("token is out of date"),
		}
	}
	nowToken := tokenPre + ":" + strconv.FormatInt(nowTs, 10)
	userPre.Token = nowToken
	err = db.UpdateCustomerToken(req.Context(), nil, userPre.CustomerId, nowToken, userPre.SessionKey)
	if err != nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken(err.Error()),
		}
	}
	return userPre, nil
}

// GenerateAccessToken creates a new access token. Returns an error if failed to generate
// random bytes.
func GenerateAccessToken() (string, error) {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	// url-safe no padding
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// ExtractAccessToken from a request, or return an error detailing what went wrong. The
// error message MUST be human-readable and comprehensible to the client.
func ExtractAccessToken(req *http.Request) (string, error) {
	//enty_logger.Info("headers are ", req.Header)
	queryToken := req.Header.Get("access_token")
	//fmt.Print("access_token is ", queryToken)
	if queryToken != "" && queryToken != "null" {
		return queryToken, nil
	}
	return "", fmt.Errorf("missing access token")
}
