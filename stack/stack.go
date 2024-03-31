package stack

type Stack interface{
	Len() int
	IsEmpty() bool
	Push(x interface{})
	Pop() (interface{}, error)
}