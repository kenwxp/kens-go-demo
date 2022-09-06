package app

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"net/http"
	"strings"
)

func Login(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := LoginReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	// first get account
	client, err := db.SelectClientByPhone(ctx, nil, reqParams.Phone)
	if err != nil || client == nil {
		return util.CommonResponse(util.CodeErr, "用户名不存在")
	}
	salt := client.Salt
	passInDB := client.Password
	pass := util.Get256Pw(reqParams.Password + salt)
	if pass != passInDB {
		return util.CommonResponse(util.CodeErr, "用户名或密码错误")
	}
	token := util.RandomString(64) + ":" + util.TimeNowUnixStr()
	err = db.UpdateClientToken(ctx, nil, client.ClientId, token)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "sql execute error")
	}
	result := LoginRep{
		Data: ClientInfoData{
			ClientId:    client.ClientId,
			ClientName:  client.ClientName,
			ClientPhone: client.ClientPhone,
		},
		Token: strings.Split(token, ":")[0],
	}
	result.Code = util.CodeOk
	result.Msg = util.MsgOk
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: result,
	}
}
func ChangePassword(
	req *http.Request, db *storage.Database, client *types.Client,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := ChangePasswordReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	salt := client.Salt
	passInDB := client.Password
	pass := util.Get256Pw(reqParams.PasswordOld + salt)
	if pass != passInDB {
		return util.CommonResponse(util.CodeErr, "原密码错误")
	}
	salt = util.RandomString(6)
	newPassword := util.Get256Pw(reqParams.PasswordNew + salt)

	err = db.UpdateClientPassword(ctx, nil, client.ClientId, salt, newPassword)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "sql execute error")
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}

func Register(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := RegisterReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	// first get account
	client, err := db.SelectClientByPhone(ctx, nil, reqParams.Phone)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	if client != nil {
		return util.CommonResponse(util.CodeErr, "用户名已存在")
	}
	salt := util.RandomString(6)
	newPassword := util.Get256Pw(reqParams.Password + salt)

	clientId, err := db.InsertClient(ctx, nil, &types.Client{
		ClientName:  reqParams.ClientName,
		ClientPhone: reqParams.Phone,
		Password:    newPassword,
		Salt:        salt,
		Token:       "",
		CreateTime:  util.TimeNowUnixStr(),
		UpdateTime:  util.TimeNowUnixStr(),
		LogonTime:   "",
		IsValid:     0,
	})
	if err != nil {
		return util.CommonResponse(util.CodeErr, "sql execute error")
	}
	println("new client_id:", clientId)
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}
