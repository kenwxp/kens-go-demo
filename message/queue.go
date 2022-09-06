package message

import (
	"sync"
)

// 代表每一个节点
type node struct {
	data *Message
	next *node
}

type queue struct {
	// 头节点
	head *node

	// 队尾节点
	rear *node

	size int

	sync.Mutex
}

func newQueue() *queue {
	q := new(queue)
	q.head = nil
	q.rear = nil
	q.size = 0
	return q
}

// Put 尾插法
func (q *queue) Put(element *Message) {
	n := new(node)
	n.data = element
	q.Lock()
	defer q.Unlock()

	if q.rear == nil {
		q.head = n
		q.rear = n
	} else {
		q.rear.next = n
		q.rear = n
	}
	q.size++
}

// PutHead 头插法，在队列头部插入一个元素
func (q *queue) PutHead(element *Message) {
	n := new(node)
	n.data = element
	q.Lock()
	defer q.Unlock()
	if q.head == nil {
		q.head = n
		q.rear = n
	} else {
		n.next = q.head
		q.head = n
	}
	q.size++
}

// Get 获取并删除队列头部的元素
func (q *queue) Get() *Message {
	if q.head == nil {
		return nil
	}
	n := q.head
	q.Lock()
	defer q.Unlock()
	// 代表队列中仅一个元素
	if n.next == nil {
		q.head = nil
		q.rear = nil

	} else {
		q.head = n.next
	}
	q.size--
	return n.data
}
