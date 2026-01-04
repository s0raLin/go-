package ArrayList

import "errors"

type Iterator interface {
	HashNext() bool     // 是否有下一个
	Next() (any, error) //下一个
	GetIndex() int
	Remove() error
	Update(newValue any)
}

type Iterable interface {
	Iterator() Iterator // 初始化接口
}

type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前索引
}

func (it *ArrayListIterator) Update(newValue any) {
	it.list.dataStore[it.currentIndex] = newValue
}

func (it *ArrayListIterator) HashNext() bool {
	return it.currentIndex < it.list.theSize
} // 是否有下一个

func (it *ArrayListIterator) Next() (any, error) {
	if !it.HashNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return value, err
} //下一个

func (it *ArrayListIterator) Remove() error {
	it.currentIndex--
	err := it.list.Delete(it.currentIndex)
	return err
}

func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}
func (list *ArrayList) Iterator() Iterator {
	iter := new(ArrayListIterator)
	iter.currentIndex = 0
	iter.list = list
	return iter
}
