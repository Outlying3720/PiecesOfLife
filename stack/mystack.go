package stack

import "errors"

type MyStack struct {
	data []interface{}
}

func (s *MyStack) Cap() int {
	return cap(s.data)
} 

func (s *MyStack) Len() int {
	return len(s.data)
} 

func (s *MyStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *MyStack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *MyStack) Pop() (interface {}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	n := len(s.data) - 1
	top := s.data[n]
	s.data[n] = nil
	s.data = s.data[:n]
	return top, nil
}

func (s *MyStack) Top() (interface {}) {
	if s.IsEmpty() {
		return nil
	}
	top := s.data[len(s.data) - 1]
	return top
}