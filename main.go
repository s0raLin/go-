package main

import (
	"first/ArrayList"
	"fmt"
)

func main() {
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
