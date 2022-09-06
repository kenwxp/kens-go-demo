package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"net/http"
	"strconv"
)

func GetRoleList(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()

	roleList := make([]RoleListOutputData, 0)
	dbRoleList, err := db.GetRoleList(ctx, nil)

	for _, roleInfo := range dbRoleList {
		var item RoleListOutputData
		item = RoleListOutputData{
			RoleId:     strconv.FormatInt(roleInfo.RoleId, 10),
			RoleName:   roleInfo.RoleName,
			MenuItems:  roleInfo.MenuItems,
			CreateTime: roleInfo.CreateTime,
			UpdateTime: roleInfo.UpdateTime,
			IsValid:    roleInfo.IsValid,
		}
		roleList = append(roleList, item)
	}

	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}

	result := GetRoleListRep{
		DataLen: int64(len(roleList)),
		Data:    roleList,
	}
	result.Code = util.CodeOk
	result.Msg = util.MsgOk
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: result,
	}
}

func AddRole(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := AddRoleReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}

	updateTs := util.TimeNowUnixStr()
	roleId, err := db.InsertRole(ctx, nil, &types.Role{
		RoleName:   reqParams.RoleName,
		MenuItems:  reqParams.MenuItems,
		CreateTime: updateTs,
		UpdateTime: updateTs,
		IsValid:    "0",
	})
	if err != nil {
		return util.CommonResponse(util.CodeErr, "sql execute error")
	}
	println("new role_id:", roleId)
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}

func DeleteRole(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := RoleIdReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	//judge
	RoleId, _ := strconv.ParseInt(reqParams.RoleId, 10, 64)
	dbRoleList, err := db.GetRoleList(ctx, nil)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	var Judge int64
	for _, roleInfo := range dbRoleList {
		if RoleId == roleInfo.RoleId {
			Judge = 1
		}
	}
	if Judge == 0 {
		return util.CommonResponse(util.CodeErr, "角色不存在")
	}

	err = db.DeleteRole(ctx, nil, RoleId)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}

func EditRole(
	req *http.Request, db *storage.Database, user *types.User,
) util.JSONResponse {
	ctx := context.TODO()
	bodyIo := req.Body
	body, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := EditRoleReq{}
	err = json.Unmarshal(body, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}

	RoleId, _ := strconv.ParseInt(reqParams.RoleId, 10, 64)
	err = db.EditRole(ctx, nil, &types.Role{
		RoleId:    RoleId,
		RoleName:  reqParams.RoleName,
		MenuItems: reqParams.MenuItems,
	})
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	return util.CommonResponse(util.CodeOk, util.MsgOk)
}
