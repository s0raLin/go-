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
	list.Insert(1, 100)
	fmt.Println(list)
}
