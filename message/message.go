package message

import (
	"context"
	"fmt"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"kens/demo/util/dict"
)

// index
func PushMessageIndex(userId string, channel dict.Channel, index int64) {
	messageKey := genMessageIndexKey(userId, channel)
	putIndex(messageKey, index)
	return
}

func GetMessageIndex(userId string, channel dict.Channel) int64 {
	messageKey := genMessageIndexKey(userId, channel)
	return readIndex(messageKey)
}

func genMessageIndexKey(userId string, channel dict.Channel) string {
	str := fmt.Sprint(userId, ":", channel)
	return util.Get256Pw(str)
}

// queue
type Message struct {
	Channel    dict.Channel
	UserId     string
	Type       string
	Topic      string
	Content    string
	KeyStr     string // 用于存储业务主键或序列号
	CreateTime string // 用于存储业务主键或序列号
}

func PushMessageInQueue(userId string, channel dict.Channel, messageType string, topic string, content string, keyStr string) error {
	messageKey := genMessageQueueKey(userId, channel, messageType)
	message := Message{
		Channel:    channel,
		UserId:     userId,
		Type:       messageType,
		Topic:      topic,
		Content:    content,
		KeyStr:     keyStr,
		CreateTime: util.TimeNowUnixStr(),
	}
	q := readQueue(messageKey)
	if q == nil {
		q = newQueue()
	}
	q.Put(&message)
	putQueue(messageKey, q)
	return nil
}

func GetOneMessageFromQueue(userId string, channel dict.Channel, messageType string) *Message {
	messageKey := genMessageQueueKey(userId, channel, messageType)
	q := readQueue(messageKey)
	if q == nil {
		return nil
	}
	return q.Get()
}

func GetAllMessageFromQueue(userId string, channel dict.Channel, messageType string) []*Message {
	messageKey := genMessageQueueKey(userId, channel, messageType)
	q := readQueue(messageKey)
	messageList := make([]*Message, 0)
	if q == nil {
		return messageList
	}
	for q.size > 0 {
		messageList = append(messageList, q.Get())
	}
	return messageList
}

func genMessageQueueKey(userId string, channel dict.Channel, messageType string) string {
	str := fmt.Sprint(userId, ":", channel, ":", messageType)
	return util.Get256Pw(str)
}

func PushMessageToDb(ctx context.Context, db *storage.Database, message types.Message) (err error) {
	message.CreateTime = util.TimeNowUnixStr()
	messageId, err := db.InsertMessage(ctx, nil, &message)
	if err != nil {
		return
	}
	PushMessageIndex(message.UserId, dict.Channel(message.MessageChannel), messageId)
	return err
}
