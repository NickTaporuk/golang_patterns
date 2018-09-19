package iterator

//var sum int = 0
//it := NewIntStatefulIterator(int_data)
//for it.Next() {
//    sum += it.Value()
//}

type intStatefulIterator struct {
	current int
	data    []int
}

func (it *intStatefulIterator) Value() int {
	return it.data[it.current]
}

func (it *intStatefulIterator) Next() bool {
	it.current++
	if it.current >= len(it.data) {
		return false
	}
	return true
}

func NewIntStatefulIterator(data []int) *intStatefulIterator {
	return &intStatefulIterator{data: data, current: -1}
}
