package main

import (
	"fmt"
	"kens/demo/util"
	"testing"
)

func TestRsa(t *testing.T) {
	data, err := util.RsaEncrypt("1234321123") //RSA加密
	fmt.Println("RSA加密", data, err)
	origData, err := util.RsaDecrypt(data) //RSA解密
	fmt.Println("RSA解密", origData, err)

}
