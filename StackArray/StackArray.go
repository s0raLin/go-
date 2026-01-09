package StackArray

import (
	"errors"
	"fmt"
)

type StackArray interface {
	Clear()            //清空
	Size() int         //大小
	Pop() (any, error) // 弹出
	Push(data any)     //压入
	IsFull() bool      //是否为满
	IsEmpty() bool     //是否为空
	String() string
}

type Stack struct {
	dataSource  []any
	capSize     int // 容量
	currentSize int // 实际使用大小
}

func (list *Stack) Clear() {
	list.dataSource = make([]any, 0, 10)
	list.capSize = 10
	list.currentSize = 0
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.dataSource = make([]any, 0, 10)
	stack.capSize = 10
	stack.currentSize = 0
	return stack
}

func (list *Stack) IsEmpty() bool {
	return list.currentSize == 0
}

func (list *Stack) IsFull() bool {
	return list.currentSize == list.capSize
}

func (list *Stack) Push(data any) {
	list.dataSource = append(list.dataSource, data)
	list.currentSize++
}

func (list *Stack) Pop() (any, error) {
	if list.IsEmpty() {
		return nil, errors.New("栈为空")
	}
	last := list.dataSource[list.currentSize-1]
	list.dataSource = list.dataSource[:list.currentSize-1]
	list.currentSize--
	return last, nil
}

func (list *Stack) Size() int {
	return list.capSize
}

func (list *Stack) String() string {
	return fmt.Sprint(list.dataSource)
}
