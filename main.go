package main

import (
	"first/ArrayList"
	"first/StackArray"
	"fmt"
)

func work1() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.Insert(1, 100)
	if err != nil {
		fmt.Println(err)
	}
	if err := list.Delete(1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}

func work2() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	for iter := list.Iterator(); iter.HashNext(); {
		next, err := iter.Next()
		if err != nil {
			break
		}
		if next == 2 {
			iter.Remove()
		}
		fmt.Println(next)
	}

	fmt.Println(list)
}

func work3() {
	stack := StackArray.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	stack.Pop()
	fmt.Println(stack)
}

func work4() {
	stack := ArrayList.NewArrayListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
	stack.Pop()
	stack.Pop()
	fmt.Println(stack)
}

func Add(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + Add(num-1)
	}
}

func work5() {
	res := Add(5)
	fmt.Println(res)
}

// 栈模拟递归
func work6() {
	mystack := StackArray.NewStack()
	mystack.Push(5)
	last := 0

	for !mystack.IsEmpty() {
		data, err := mystack.Pop()
		if err != nil {
			break
		}
		if data == 0 {
			last += 0
		} else {
			last += data.(int)
			mystack.Push(data.(int) - 1)
		}
	}
	fmt.Println(last)
}

// 斐波那契数列
func FAB(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return FAB(num-1) + FAB(num-2)
	}
}

func work7() {
	println(FAB(5))
	mystack := StackArray.NewStack()
	mystack.Push(5)
	last := 0
	for !mystack.IsEmpty() {
		data, err := mystack.Pop()
		if err != nil {
			break
		}
		if data == 1 || data == 2 {
			last += 1
		} else {
			mystack.Push(data.(int) - 1)
			mystack.Push(data.(int) - 2)
		}
	}
	fmt.Println(last)
}

func main() {
	work7()
}
