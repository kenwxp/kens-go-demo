package demo

import (
	"fmt"
	"github.com/gorilla/mux"
	"kens/demo/httputil"
	"kens/demo/storage"
	"net/http"
)

// Setup configures the given mux with sync-server listeners
func InitRouting(
	router *mux.Router, db *storage.Database, wsCli *httputil.WebSocketCli,
) {

	subRouter := router.PathPrefix("/demo").Subrouter()
	// upgrade websocket 创建路由
	subRouter.HandleFunc("/socket/message", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("api /socket/message...")
		//account, err := auth.VerifyUserFromRequest(r, db)
		//if err != nil || account == nil {
		//	conn, _ := socket.NewUpgrader(r).Upgrade(w, r, nil)
		//	_ = conn.WriteMessage(websocket.TextMessage, []byte("missing token"))
		//	_ = conn.Close()
		//	return
		//}
		wsCli.NewWebSocketConn(w, r, "1", httputil.AppChannel, "message")
		wsCli.Send("1", httputil.AppChannel, "message", []byte("已连接"))
	})

	subRouter.HandleFunc("/add/message", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("api /add/message...")
		//account, err := auth.VerifyUserFromRequest(r, db)
		//if err != nil || account == nil {
		//	conn, _ := socket.NewUpgrader(r).Upgrade(w, r, nil)
		//	_ = conn.WriteMessage(websocket.TextMessage, []byte("missing token"))
		//	_ = conn.Close()
		//	return
		//}
		wsCli.Send("1", httputil.AppChannel, "message", []byte("hello world"))
	})
}
