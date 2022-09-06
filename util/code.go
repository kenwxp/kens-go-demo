package util

import "net/http"

const (
	CodeOk     = "0"  // 成功
	CodeErr    = "-1" // 失败
	CodeEmpty  = "1"  // 无或空
	CodeReject = "2"  // 拒绝
	MsgOk      = "操作成功"
	MsgErr     = "操作失败"
	MsgEmpty   = "空或无"
	MsgReject  = "操作拒绝"
)

type CommonRep struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func CommonResponse(retCode string, retMsg string) JSONResponse {
	return JSONResponse{
		Code: http.StatusOK,
		JSON: CommonRep{
			Code: retCode,
			Msg:  retMsg,
		},
	}
}
