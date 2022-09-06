package message

import (
	"sync"
)

var messageQueueTagMap = func() map[string]*queue {
	return make(map[string]*queue)
}()

var QueueLock = func() *sync.Mutex {
	return &sync.Mutex{}
}()

func putQueue(tag string, val *queue) {
	QueueLock.Lock()
	messageQueueTagMap[tag] = val
	QueueLock.Unlock()
}

// read and del
func readThenDelQueue(tag string) *queue {
	QueueLock.Lock()
	defer QueueLock.Unlock()
	res := messageQueueTagMap[tag]
	delete(messageQueueTagMap, tag)
	return res
}

// read and del
func deleteQueue(tag string) {
	QueueLock.Lock()
	defer QueueLock.Unlock()
	delete(messageQueueTagMap, tag)
}

// read and del
func readQueue(tag string) *queue {
	QueueLock.Lock()
	defer QueueLock.Unlock()
	res := messageQueueTagMap[tag]
	return res
}
