package uni

type CommonHttpRep struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type CheckNewCustomerReq struct {
	LoginCode string `json:"loginCode"` // 微信登录code
}

type CheckNewCustomerRep struct {
	IsNew string `json:"isNew"`
	CommonHttpRep
}
type CustomerLoginReq struct {
	LoginCode string `json:"loginCode"` // 微信登录code
	PhoneCode string `json:"phoneCode"` // 手机号获取code
	NickName  string `json:"nickName"`  // 昵称
	Avatar    string `json:"avatar"`    // 头像
}

type CustomerLoginRep struct {
	Data CustomerInfo `json:"data"`
	CommonHttpRep
}

type CustomerInfo struct {
	OpenId      string `json:"openId "`     //	用户openId
	AccessToken string `json:"accessToken"` //	后台登录toke
}

type OpenDoorReq struct {
	ShopId string `json:"shopId"`
}
