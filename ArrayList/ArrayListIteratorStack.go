package ArrayList

import (
	"fmt"
)

type StackX struct {
	array *ArrayList
	iter  Iterator
}

func NewArrayListIteratorStack() *StackX {
	myStackX := new(StackX)
	myStackX.array = NewArrayList()
	myStackX.iter = myStackX.array.Iterator()
	return myStackX
}

func (list *StackX) Clear() {
	list.array.Clear()
}

func (list *StackX) IsEmpty() bool {
	return list.array.theSize == 0
}

func (list *StackX) Append(data any) {
	list.array.Append(data)
}

func (list *StackX) Size() int {
	return list.array.Size()
}

func (list *StackX) String() string {
	return fmt.Sprint(list.array)
}
