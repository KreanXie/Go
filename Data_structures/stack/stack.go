package stack

import (
	"fmt"
	"math"
)

const inf = math.MaxInt64

type node struct {
	val        int
	last, next *node
}

type Stack struct {
	head, tail *node
}

func (s *Stack) Empty() bool {
	return s.head == nil
}

func (s *Stack) Push(val int) {
	if s.head == nil {
		s.head = &node{val, s.head, nil}
		s.tail = s.head
	} else {
		s.tail.next = &node{val, s.tail, nil}
		s.tail = s.tail.next
	}
}

func (s *Stack) Pop() {
	if s.Empty() {
		return
	} else {
		if s.tail == s.head {
			s.tail, s.head = nil, nil
		} else {
			s.tail = s.tail.last
			s.tail.next = nil
		}
	}
}

func (s *Stack) Top() int {
	if s.Empty() {
		return inf
	}
	return s.tail.val
}

func (s *Stack) Clear() {
	s.head, s.tail = nil, nil
}

func (s Stack) Print() {
	fmt.Printf("[")
	cur := s.head
	for cur != nil {
		fmt.Printf("%d ", cur.val)
		cur = cur.next
	}
	fmt.Printf("\b]\n")
}
