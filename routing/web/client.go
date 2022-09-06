package web

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"kens/demo/storage"
	"kens/demo/storage/types"
	_ "kens/demo/storage/types"
	"kens/demo/util"
	"kens/demo/util/dict"
	"net/http"
	_ "strconv"
)

func GetClientListWithCondition(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	reqParams := GetClientListReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin && dict.UserRole(user.RoleId) != dict.RoleApproveUser {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}
	clientList := make([]ClientListOutputData, 0)
	dbClientList, err := db.SelectClientListWithCondition(ctx, nil, reqParams.ClientName, reqParams.Phone)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	for _, item := range dbClientList {
		clientList = append(clientList, ClientListOutputData{
			ClientId:   item.ClientId,
			ClientName: item.ClientName,
			Phone:      item.ClientPhone,
			CreateTime: util.TimeStampBeautify(item.CreateTime),
			UpdateTime: util.TimeStampBeautify(item.UpdateTime),
			LogonTime:  util.TimeStampBeautify(item.LogonTime),
			IsValid:    item.IsValid,
		})
	}
	result := GetClientListRep{
		DataLen: int64(len(clientList)),
		Data:    clientList,
	}
	result.Code = util.CodeOk
	result.Msg = util.MsgOk
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: result,
	}
}

func AddClient(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	println(string(body))
	reqParams := AddClientReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin && dict.UserRole(user.RoleId) != dict.RoleApproveUser {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}
	if reqParams.ClientName == "" {
		return util.CommonResponse(util.CodeErr, "用户名不能为空")
	}
	if !util.VerifyMobileFormat(reqParams.Phone) {
		return util.CommonResponse(util.CodeErr, "手机号非法")
	}
	// first get account
	err = db.WithTransaction(func(txn *sql.Tx) error {
		client, err := db.SelectClientByPhone(ctx, txn, reqParams.Phone)
		if err != nil {
			return err
		}
		if client != nil {
			return errors.New("手机号已存在")
		}
		// 设定手机后六位为初始密码
		salt := util.RandomString(6)
		phone := reqParams.Phone
		password := phone[len(phone)-6:]
		newPassword := util.Get256Pw(password + salt)
		_, err = db.InsertClient(ctx, txn, &types.Client{
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
			return err
		}

		return nil
	})
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}

	return util.CommonResponse(util.CodeOk, util.MsgOk)
}

func DeleteClient(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := DeleteClientReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin && dict.UserRole(user.RoleId) != dict.RoleApproveUser {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}
	err = db.DeleteClientByClientId(ctx, nil, reqParams.ClientId)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}

func EditClient(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := EditClientReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin && dict.UserRole(user.RoleId) != dict.RoleApproveUser {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}
	if reqParams.ClientName == "" {
		return util.CommonResponse(util.CodeErr, "用户名不能为空")
	}
	if !util.VerifyMobileFormat(reqParams.Phone) {
		return util.CommonResponse(util.CodeErr, "手机号非法")
	}
	err = db.EditClient(ctx, nil, reqParams.ClientId, reqParams.ClientName, reqParams.Phone)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}

func ResetClientPassword(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := ResetClientPasswordReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin && dict.UserRole(user.RoleId) != dict.RoleApproveUser {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}
	err = db.WithTransaction(func(txn *sql.Tx) error {
		client, err1 := db.SelectClientByClientId(ctx, txn, reqParams.ClientId)
		if err1 != nil {
			return err1
		}
		if client == nil {
			return errors.New("客户不存在")
		}
		phone := client.ClientPhone
		// 重设密码
		salt := util.RandomString(6)
		password := phone[len(phone)-6:]
		newPassword := util.Get256Pw(password + salt)
		err1 = db.UpdateClientPassword(ctx, txn, reqParams.ClientId, salt, newPassword)
		if err1 != nil {
			return err1
		}
		return nil
	})

	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}
