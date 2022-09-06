package uni

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"kens/demo/pkg/wxapi"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"net/http"
)

func CheckNewCustomer(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	ctx := req.Context()
	bodyIo := req.Body
	reqBody, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := CheckNewCustomerReq{}
	err = json.Unmarshal(reqBody, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	userSession, err := wxapi.GetUserSession(ctx, reqParams.LoginCode)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	customer, err := db.SelectCustomerByOpenId(ctx, nil, userSession.Openid)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "sql execute error")
	}
	isNew := "0"
	if customer == nil {
		isNew = "1"
	}
	result := CheckNewCustomerRep{
		IsNew: isNew,
	}
	result.Code = util.CodeOk
	result.Msg = util.MsgOk
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: result,
	}
}

func CustomerLogin(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	ctx := req.Context()
	bodyIo := req.Body
	reqBody, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "read body err")
	}
	reqParams := CustomerLoginReq{}
	err = json.Unmarshal(reqBody, &reqParams)
	if err != nil {
		return util.CommonResponse(util.CodeErr, "Unmarshal json err")
	}
	userSession, err := wxapi.GetUserSession(ctx, reqParams.LoginCode)
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	result := CustomerLoginRep{}
	err = db.WithTransaction(func(txn *sql.Tx) error {
		customer, err := db.SelectCustomerByOpenId(ctx, txn, userSession.Openid)
		if err != nil {
			return err
		}
		timeNow := util.TimeNowUnixStr()
		token := util.RandomString(64)
		tokenStr := token + ":" + timeNow

		if customer == nil {
			phone, err := wxapi.GetUserPhone(ctx, reqParams.PhoneCode)
			if err != nil {
				return err
			}
			newCustomer := types.Customer{
				Nickname:    reqParams.NickName,
				Avatar:      reqParams.Avatar,
				Phone:       phone,
				Gender:      "0",
				Birth:       "1990-01-01",
				Token:       tokenStr,
				OpenId:      userSession.Openid,     // 微信openid
				SessionKey:  userSession.SessionKey, // session_key
				IsMember:    0,
				JoinTime:    "",
				JoinEndTime: "",
				CreateTime:  timeNow,
				UpdateTime:  timeNow,
				IsBlack:     0,
			}
			_, err = db.InsertCustomer(ctx, txn, &newCustomer)
			if err != nil {
				return err
			}
			result = CustomerLoginRep{
				Data: CustomerInfo{
					OpenId:      userSession.Openid, //	用户openId
					AccessToken: token,              //	后台登录toke
				},
			}
		} else {
			err = db.UpdateCustomerToken(ctx, txn, customer.CustomerId, tokenStr, userSession.SessionKey)
			if err != nil {
				return err
			}
			result = CustomerLoginRep{
				Data: CustomerInfo{
					OpenId:      userSession.Openid, //	用户openId
					AccessToken: token,              //	后台登录toke
				},
			}
		}
		return nil
	})
	if err != nil {
		return util.CommonResponse(util.CodeErr, err.Error())
	}
	result.Code = util.CodeOk
	result.Msg = util.MsgOk
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: result,
	}
}
