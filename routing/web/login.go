package web

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"net/http"
	"strconv"
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
	user, err := db.SelectUserByAccNo(ctx, nil, reqParams.AccNo)
	if err != nil || user == nil {
		return util.CommonResponse(util.CodeErr, "用户名不存在")
	}
	salt := user.Salt
	passInDB := user.Password
	pass := util.Get256Pw(reqParams.Password + salt)
	if pass != passInDB {
		return util.CommonResponse(util.CodeErr, "用户名或密码错误")
	}
	token := util.RandomString(64) + ":" + util.TimeNowUnixStr()
	err = db.WithTransaction(func(txn *sql.Tx) error {
		err = db.UpdateUserToken(ctx, txn, user.AccNo, token)
		if err != nil {
			return err
		}
		err = db.UpdateUserLogon(ctx, txn, user.AccNo, "1")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Marshal json err")
	}
	userInfo := UserInfoData{}
	err = json.Unmarshal(jsonBytes, &userInfo)
	userInfo.RoleId = strconv.FormatInt(user.RoleId, 10)
	userInfo.LogonTime = util.TimeStampBeautify(util.TimeNowUnixStr())
	result := LoginRep{
		UserInfo: userInfo,
		Token:    strings.Split(token, ":")[0],
	}
	result.Code = util.CodeOk
	result.Msg = util.MsgOk
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: result,
	}
}
func ChangePassword(
	req *http.Request, db *storage.Database, user *types.User,
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
	salt := user.Salt
	passInDB := user.Password
	pass := util.Get256Pw(reqParams.PasswordOld + salt)
	if pass != passInDB {
		return util.CommonResponse(util.CodeErr, "原密码错误")
	}
	salt = util.RandomString(6)
	newPassword := util.Get256Pw(reqParams.PasswordNew + salt)

	err = db.UpdateUserPassword(ctx, nil, user.AccNo, salt, newPassword)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "sql execute error")
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}
