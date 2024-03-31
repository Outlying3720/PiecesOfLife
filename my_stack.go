package main

import "fmt"
import "runtime"

import "errors"

type Stack interface{
	Len() int
	IsEmpty() bool
	Push(x interface{})
	Pop() (interface{}, error)
}

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


func main() {
	var IStack Stack
	var myStack = new(MyStack)
	IStack = myStack
	fmt.Println(IStack)
	fmt.Println(IStack.IsEmpty())
	fmt.Println(IStack.Len())
	IStack.Push(6)
	IStack.Push("String")
	IStack.Push(3.1415)
	fmt.Println(IStack)
	fmt.Println(IStack.IsEmpty())
	fmt.Println(IStack.Len())
	fmt.Println(IStack.Pop())
	fmt.Println(IStack.Pop())
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	fmt.Println(myStack.Cap())

	for i := 0; i < 10000; i++ {
		IStack.Push(3.1415)
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	fmt.Println(myStack.Cap())

	// runtime.GC()

	// runtime.ReadMemStats(&m)
	// fmt.Printf("%d Kb\n", m.Alloc / 1024)
	// fmt.Println(myStack.Cap())

	for i := 0; i < 10000; i++ {
		IStack.Pop()
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	fmt.Println(myStack.Cap())

	runtime.GC()

	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	fmt.Println(myStack.Cap())
}