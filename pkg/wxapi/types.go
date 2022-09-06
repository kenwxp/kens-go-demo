package wxapi

type CommonHttpRep struct {
	ErrCode int64  `json:"errcode"` // 错误码
	ErrMsg  string `json:"errmsg"`  // 错误信息
}

type AccessToken struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int64  `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	CommonHttpRep
}

type UserSession struct {
	Openid     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionId    string `json:"unionid"`     // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
	CommonHttpRep
}
type UserPhoneInfoReq struct {
	Code string `json:"code"` // code
}

type UserPhoneInfoRep struct {
	PhoneInfo PhoneInfo `json:"phone_info"` // 用户手机号信息
	CommonHttpRep
}

type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string    `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string    `json:"countryCode"`     // 区号
	WaterMark       WaterMark `json:"watermark"`       // 数据水印
}
type WaterMark struct {
	AppId     string `json:"appid"`     // 小程序appid
	Timestamp int64  `json:"timestamp"` // 用户获取手机号操作的时间戳
}
