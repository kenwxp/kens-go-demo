package httputil

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"kens/demo/util"
	"net/http"
	"sync"
	"time"
)

type Channel string

const (
	AppChannel     = Channel("0")
	MiniAppChannel = Channel("1")
	WindowsChannel = Channel("2")
	DemoChannel    = Channel("3")
)

type WebSocketCli struct {
	mu    sync.Mutex
	conns map[Channel]map[string]map[string]*websocket.Conn
}

//type WebSocketConn struct {
//	Channel Channel
//	UserId  string
//	SubType string
//	Conn    *websocket.Conn
//}

func NewWebsocketCli() *WebSocketCli {
	wc := &WebSocketCli{
		mu: sync.Mutex{},
	}
	return wc
}

func (wc *WebSocketCli) NewWebSocketConn(w http.ResponseWriter, r *http.Request, userId string, channel Channel, subType string) {
	var conn *websocket.Conn
	var err error
	//userId := user.ID
	defer func() {
		fmt.Println("defer 断开连接，渠道：", channel, "用户Id：", userId, "处理类型：", subType, util.TimeNowFormatString())
		wc.mu.Lock()
		if conn != nil {
			_ = conn.Close()
		}
		delete(wc.conns[channel][userId], subType)
		wc.mu.Unlock()
	}()

	// drop the connection right away if exist
	wc.mu.Lock()
	if _, ok := wc.conns[channel][userId][subType]; ok {
		fmt.Println("websocket 连接已存在 重复连接！")
		wc.mu.Unlock()
		return
	}
	// make new connection
	conn, err = NewUpGrader(r).Upgrade(w, r, nil)
	if err != nil {
		wc.mu.Unlock()
		return
	}
	// set conn
	wc.conns = putDeeply(wc.conns, userId, channel, subType, conn)
	wc.mu.Unlock()
	fmt.Println("websocket 连接已创建，渠道：", channel, "用户Id：", userId, "处理类型：", subType, util.TimeNowFormatString())

	conn.SetPingHandler(func(appData string) error {
		fmt.Println("ping:", appData)
		_ = conn.WriteMessage(websocket.PongMessage, []byte(appData))
		return conn.SetReadDeadline(time.Now().Add(time.Second * 3))
	})
	conn.SetPongHandler(func(s string) error {
		fmt.Println("ping:", s)
		_ = conn.WriteMessage(websocket.PingMessage, []byte(s))
		return conn.SetReadDeadline(time.Now().Add(time.Second * 3))
	})
	for {
		_, _, err = conn.NextReader()
		fmt.Println(err)
		if err != nil {
			break
		}
	}
}

func putDeeply(sourceMap interface{}, userId string, channel Channel, subType string, conn *websocket.Conn) map[Channel]map[string]map[string]*websocket.Conn {
	if sourceMap != nil {
		srcMap := sourceMap.(map[Channel]map[string]map[string]*websocket.Conn)
		if _, ok := srcMap[channel]; ok {
			if _, ok = srcMap[channel][userId]; ok {
				srcMap[channel][userId][subType] = conn
				return srcMap
			} else {
				subTypeMap := make(map[string]*websocket.Conn)
				subTypeMap[subType] = conn
				srcMap[channel][userId] = subTypeMap
				return srcMap
			}
		}
	}
	subTypeMap := make(map[string]*websocket.Conn)
	subTypeMap[subType] = conn
	userIdMap := make(map[string]map[string]*websocket.Conn)
	userIdMap[userId] = subTypeMap
	channelMap := make(map[Channel]map[string]map[string]*websocket.Conn)
	channelMap[channel] = userIdMap
	return channelMap
}

func (wc *WebSocketCli) Send(userId string, channel Channel, subType string, data []byte) error {
	var conn *websocket.Conn
	wc.mu.Lock()
	if c, ok := wc.conns[channel][userId][subType]; ok {
		conn = c
		wc.mu.Unlock()
	} else {
		wc.mu.Unlock()
		return errors.New("websocket 连接未创建，渠道：" + string(channel) + "， 用户Id：" + userId + "， 处理类型：" + subType + "， 时间：" + util.TimeNowFormatString())
	}
	if conn != nil {
		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			return errors.New("websocket 发送错误，渠道：" + string(channel) + "， 用户Id：" + userId + "， 处理类型：" + subType + "， 时间：" + util.TimeNowFormatString())
		}
	}
	return nil
}

func NewUpGrader(r *http.Request) *websocket.Upgrader {
	return &websocket.Upgrader{
		EnableCompression: true,
		Subprotocols:      []string{r.Header.Get("Sec-WebSocket-Protocol")},
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
