package demo

import (
	"kens/demo/httputil"
	"kens/demo/storage"
	"net/http"
)

func AddMessage(
	req *http.Request, db *storage.Database, wsCli *httputil.WebSocketCli,
) {
	wsCli.Send("1", httputil.AppChannel, "message", []byte("hello world"))
}
