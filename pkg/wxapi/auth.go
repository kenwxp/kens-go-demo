package wxapi

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"kens/demo/cache"
	"kens/demo/util/enty_logger"
	"net/http"
	"time"
)

const (
	appid         = "wxdee2b19db3be8517"
	secret        = "99887628a67e7d8a23d4d8410bcf47a1"
	wxAccessToken = "wx_access_token"
)

// 获取小程序全局唯一后台接口调用凭据（access_token）
func GetAccessToken(ctx context.Context) (*AccessToken, error) {
	// GET https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	getAccessTokenUri := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getAccessTokenUri, nil)

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
	enty_logger.Debug(string(body))
	rep := &AccessToken{}
	err = json.Unmarshal(body, rep)
	if err != nil {
		return nil, err
	}
	return rep, nil
}

func setWxHttpHeader(req *http.Request) error {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Close = true
	return nil
}

func wxAuthentication(ctx context.Context) (string, error) {
	accessToken, isExist := cache.CacheAll.Get(wxAccessToken)
	if isExist {
		return accessToken.(string), nil
	}
	rep, err := GetAccessToken(ctx)
	if err != nil {
		return "", err
	}
	if rep.ErrCode != 0 {
		return "", errors.New(rep.ErrMsg)
	}
	newAccessToken := rep.AccessToken
	newExpiresIn := rep.ExpiresIn
	cache.CacheAll.Set(wxAccessToken, newAccessToken, time.Duration(newExpiresIn)*time.Second)
	enty_logger.Info("wx access token update:", newAccessToken)
	return newAccessToken, nil
}
