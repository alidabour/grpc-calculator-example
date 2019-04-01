package stack

import "errors"

//Stack data structure
type Stack interface {
	IsEmpty() bool
	Top() string
	Pop() (Stack, error)
	Push(string) Stack
}

type stack []string

//New return new stack
func New() Stack {
	return stack{}
}

//IsEmpty check if stack has elements
func (s stack) IsEmpty() bool {
	return len(s) == 0
}

//Top return the top element
func (s stack) Top() string {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return ""
}

//Pop return a stack without the top element
func (s stack) Pop() (Stack, error) {
	if !s.IsEmpty() {
		return s[:len(s)-1], nil
	}
	return s, errors.New("Stack is empty")
}

//Push add to the top of the stack
func (s stack) Push(i string) Stack {
	return append(s, i)
}
