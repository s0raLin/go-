package ArrayList

import "errors"

type Iterator interface {
	HashNext() bool     // 是否有下一个
	Next() (any, error) //下一个
	GetIndex() int
	Remove()
}

type Iterable interface {
	Iterator() Iterator // 初始化接口
}

type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前索引
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

func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}

func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}
func (it *ArrayList) Iterator() Iterator {
	iter := new(ArrayListIterator)
	iter.currentIndex = 0
	iter.list = it
	return iter
}
