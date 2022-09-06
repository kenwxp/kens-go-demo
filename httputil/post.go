package httputil

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type RegisterReq struct {
	Secret string
}

type RegisterResp struct {
	Token string
}

// return string as token ,error
func PostRegisterEntyPay(secret string) (string, error) {
	req := RegisterReq{Secret: secret}
	url := "localhost:10001/r0/business/register"
	//url := "entypay.io/api/business/register"
	res, err := Post(url, req, "application/json")
	if err != nil {
		return "", err
	}
	tokenStruct := RegisterResp{}
	err = json.Unmarshal(res, &tokenStruct)
	if err != nil {
		return "", err
	}
	return tokenStruct.Token, nil
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) ([]byte, error) {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result, nil
}
