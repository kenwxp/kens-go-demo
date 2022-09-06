package message

import (
	"testing"
)

func TestMessage(t *testing.T) {
	PushMessage("1", AppChannel, "notice", "", "测试1", "")
	PushMessage("1", AppChannel, "notice", "", "测试2", "")
	PushMessage("1", AppChannel, "notice", "", "测试3", "")
	PushMessage("1", AppChannel, "notice", "", "测试4", "")
	PushMessage("1", AppChannel, "notice", "", "测试5", "")

	messageList := GetAllMessage("1", AppChannel, "notice")
	for _, message := range messageList {
		println(message.Content)
	}
	//println(GetOneMessage("1", AppChannel, "notice").Content)
	//println(GetOneMessage("1", AppChannel, "notice").Content)
	//println(GetOneMessage("1", AppChannel, "notice") == nil)

}
