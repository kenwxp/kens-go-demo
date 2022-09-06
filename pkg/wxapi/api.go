package wxapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"kens/demo/util/enty_logger"
	"net/http"
)

// 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程
func GetUserSession(ctx context.Context, jsCode string) (*UserSession, error) {
	// GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	getUserSessionUri := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appid + "&secret=" + secret + "&js_code=" + jsCode + "&grant_type=authorization_code"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getUserSessionUri, nil)

	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	enty_logger.Debug("GetUserSession resp:", string(body))
	rep := &UserSession{}
	err = json.Unmarshal(body, rep)

	return rep, nil
}

// code换取用户手机号。 每个 code 只能使用一次，code的有效期为5min
func GetUserPhoneInfo(ctx context.Context, code string) (*UserPhoneInfoRep, error) {
	// POST https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=ACCESS_TOKEN
	accessToken, err := wxAuthentication(ctx)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := json.Marshal(
		UserPhoneInfoReq{
			Code: code,
		})
	if err != nil {
		return nil, err
	}
	getUserPhoneUri := "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=" + accessToken
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, getUserPhoneUri, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	setWxHttpHeader(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	enty_logger.Debug("GetUserPhoneInfo resp:", string(body))
	rep := &UserPhoneInfoRep{}
	err = json.Unmarshal(body, rep)
	if err != nil {
		return nil, err
	}
	return rep, nil
}

func GetUserPhone(ctx context.Context, code string) (string, error) {
	userPhoneInfo, err := GetUserPhoneInfo(ctx, code)
	if err != nil {
		return "", err
	}
	if userPhoneInfo.ErrCode != 0 {
		return "", errors.New(userPhoneInfo.ErrMsg)
	}
	return userPhoneInfo.PhoneInfo.PurePhoneNumber, nil
}
