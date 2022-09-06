package web

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"kens/demo/util/dict"
	"net/http"
	"strconv"
)

func GetUserList(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	reqParams := &GetUserListReq{}
	err = json.Unmarshal(body, reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin && dict.UserRole(user.RoleId) != dict.RoleApproveUser {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}
	dbUserList, err := db.SelectUserListWithCondition(ctx, nil, types.User{
		AccNo:   reqParams.AccNo,
		AccName: reqParams.AccName,
		RoleId:  reqParams.RoleId,
		OnWork:  reqParams.OnWork,
	})
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	userList := make([]UserListItemData, 0)
	for _, item := range dbUserList {
		itemBytes, err := json.Marshal(item)
		if err != nil {
			return util.CommonResponse(util.CodeErr, err.Error())
		}
		var temp UserListItemData
		json.Unmarshal(itemBytes, &temp)
		temp.CreateTime = util.TimeStampBeautify(temp.CreateTime)
		temp.UpdateTime = util.TimeStampBeautify(temp.UpdateTime)
		temp.LogonTime = util.TimeStampBeautify(temp.LogonTime)
		userList = append(userList, temp)
	}

	result := GetUserListRep{
		DataLen: int64(len(userList)),
		Data:    userList,
	}
	result.Code = util.CodeOk
	result.Msg = util.MsgOk
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: result,
	}
}

func AddUser(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	reqParams := &AddUserReq{}
	err = json.Unmarshal(body, reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}
	roleId := reqParams.RoleId
	if roleId == 0 {
		return util.CommonResponse(util.CodeErr, "必须指定用户角色")
	}
	err = db.WithTransaction(func(txn *sql.Tx) error {
		maxAccNo, err := db.SelectMaxAccNoByRoleId(ctx, txn, roleId)
		if err != nil {
			return err
		}
		if maxAccNo == 0 {
			maxAccNo = roleId * 100000
		}
		// 设定手机后六位为初始密码
		salt := util.RandomString(6)
		password := reqParams.Phone[len(reqParams.Phone)-6:]
		newPassword := util.Get256Pw(password + salt)
		err = db.InsertUser(ctx, txn, &types.User{
			AccNo:      strconv.FormatInt(maxAccNo+1, 10),
			AccName:    reqParams.AccName,
			Password:   newPassword,
			Salt:       salt,
			Token:      "",
			Phone:      reqParams.Phone,
			Email:      reqParams.Email,
			RoleId:     roleId,
			OnWork:     "0",
			CreateTime: util.TimeNowUnixStr(),
			UpdateTime: util.TimeNowUnixStr(),
			LogonTime:  "",
			IsValid:    0,
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

func DeleteUser(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := AccNoReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}
	err = db.DeleteUser(ctx, nil, reqParams.AccNo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}

func EditUser(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	reqParams := &EditUserReq{}
	err = json.Unmarshal(body, reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	if dict.UserRole(user.RoleId) != dict.RoleSuperAdmin {
		return util.CommonResponse(util.CodeErr, "当前用户无操作权限")
	}

	err = db.EditUser(ctx, nil, &types.User{
		AccNo:   reqParams.AccNo,
		AccName: reqParams.AccName,
		Phone:   reqParams.Phone,
		Email:   reqParams.Email,
	})
	if err != nil {
		return util.CommonResponse(util.CodeErr, "json input is wrong")
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}

func ResetUserPassword(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := ResetUserPasswordReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}

	dbUser, err := db.SelectUserByAccNo(ctx, nil, reqParams.AccNo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	if dbUser == nil {
		return util.CommonResponse(util.CodeErr, "用户名不存在")
	}
	// 设定手机后六位为初始密码
	salt := util.RandomString(6)
	password := dbUser.Phone[len(dbUser.Phone)-6:]
	newPassword := util.Get256Pw(password + salt)
	err = db.UpdateUserPassword(ctx, nil, reqParams.AccNo, salt, newPassword)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}
