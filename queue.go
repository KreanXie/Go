package queue

import (
	"fmt"
	"math"
)

const inf = math.MaxInt64

type Queue struct {
	head, tail *node
}

type node struct {
	data int
	next *node
}

func (q *Queue) Enqueue(val int) {
	if q.head == nil {
		q.tail = &node{val, nil}
		q.head = q.tail
	} else if q.head == q.tail {
		q.tail.next = &node{val, nil}
		q.tail = q.tail.next
		q.head.next = q.tail
	} else {
		q.tail.next = &node{val, nil}
		q.tail = q.tail.next
	}
}

func (q *Queue) Dequeue() {
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
}

func (q *Queue) Front() int {
	if q.head == nil {
		return inf
	}
	return q.head.data
}

func (q *Queue) Back() int {
	if q.head == nil {
		return inf
	}
	return q.tail.data
}

func (q *Queue) Empty() bool {
	return q.head == nil
}

func (q *Queue) Clear() {
	q.head, q.tail = nil, nil
}

func (q Queue) Print() {
	fmt.Printf("[")
	cur := q.head
	for cur != nil {
		fmt.Printf("%d", cur.data)
		if cur.next != nil {
			fmt.Printf(" ")
		}
		cur = cur.next
	}
	fmt.Printf("]\n")
}
