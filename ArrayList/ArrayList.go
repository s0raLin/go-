package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (any, error)
	Set(index int, newValue any) error
	Insert(index int, newValue any) error
	Append(newValue any)
	Clear() // 清空数组
	Delete(index int) error
	String() string
}

type ArrayList struct {
	dataStore []any
	theSize   int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]any, 0, 10) // 开辟10个
	list.theSize = 0
	return list
}

func (list *ArrayList) Append(newValue any) {
	list.dataStore = append(list.dataStore, newValue)
	list.theSize++
}

func (list *ArrayList) checkFull(index int, newValue any) {
	if list.theSize != cap(list.dataStore) {
		return
	}
	newDataStore := make([]any, 2*list.theSize)
	copy(newDataStore, list.dataStore)
	newDataStore = append(newDataStore[:index], append([]any{newValue}, newDataStore[index:]...)...)
	list.dataStore = newDataStore
}

func (list *ArrayList) Insert(index int, newValue any) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}

	list.dataStore = append(list.dataStore[:index], append([]any{newValue}, list.dataStore[index:]...)...)

	return nil
}

func (list *ArrayList) Size() int {
	return list.theSize
}

func (list *ArrayList) Get(index int) (any, error) {
	if index < 0 || index >= len(list.dataStore) {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newValue any) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newValue
	return nil
}

func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	return nil
}
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]any, 0, 10)
	list.theSize = 0
}
