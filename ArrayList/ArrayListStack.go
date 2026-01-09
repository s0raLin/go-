package ArrayList

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

type ArrayListStack struct {
	array       *ArrayList
	capSize     int // 容量
	currentSize int // 实际使用大小
}

func (list *ArrayListStack) Clear() {
	list.array.Clear()
	list.capSize = 10
	list.currentSize = 0
}

func NewArrayListStack() *ArrayListStack {
	stack := new(ArrayListStack)
	stack.array = NewArrayList()
	stack.capSize = 10
	stack.currentSize = 0
	return stack
}

func (list *ArrayListStack) IsEmpty() bool {
	return list.currentSize == 0
}

func (list *ArrayListStack) IsFull() bool {
	return list.currentSize == list.capSize
}

func (list *ArrayListStack) Push(data any) {
	list.array.Append(data)
	list.currentSize++
}

func (list *ArrayListStack) Pop() (any, error) {
	if list.IsEmpty() {
		return nil, errors.New("栈为空")
	}
	last := list.array.dataStore[list.currentSize-1]
	list.array.dataStore = list.array.dataStore[:list.currentSize-1]
	list.currentSize--
	return last, nil
}

func (list *ArrayListStack) Size() int {
	return list.capSize
}

func (list *ArrayListStack) String() string {
	return fmt.Sprint(list.array)
}
