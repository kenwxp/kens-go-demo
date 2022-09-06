package message

import "sync"

var messageIndexTagMap = func() map[string]int64 {
	return make(map[string]int64)
}()

var Lock = func() *sync.Mutex {
	return &sync.Mutex{}
}()

func putIndex(tag string, val int64) {
	Lock.Lock()
	messageIndexTagMap[tag] = val
	Lock.Unlock()
}

// read and del
func readThenDelIndex(tag string) int64 {
	Lock.Lock()
	defer Lock.Unlock()
	res := messageIndexTagMap[tag]
	delete(messageIndexTagMap, tag)
	return res
}

// read and del
func deleteIndex(tag string) {
	Lock.Lock()
	defer Lock.Unlock()
	delete(messageIndexTagMap, tag)
}

// read and del
func readIndex(tag string) int64 {
	Lock.Lock()
	defer Lock.Unlock()
	res := messageIndexTagMap[tag]
	return res
}
