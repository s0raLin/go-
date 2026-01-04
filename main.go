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
	fmt.Println(stack)
}

func main() {
	work3()
}
