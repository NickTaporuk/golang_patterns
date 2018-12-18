package iterator

type StatefulIterator interface {
	Value() int
	Next() bool
}

//
type intStatefulIterator struct {
	current int
	data    []int
}

//
func (it *intStatefulIterator) Value() int {
	return it.data[it.current]
}

//
func (it *intStatefulIterator) Next() bool {
	it.current++
	if it.current >= len(it.data) {
		return false
	}
	return true
}

//
func NewIntStatefulIterator(data []int) *intStatefulIterator {
	return &intStatefulIterator{data: data, current: -1}
}

// Example :
//var sum int = 0
// int_data := []int{1,2,3,4,5,6,7,89}
//it := NewIntStatefulIterator(int_data)
//for it.Next() {
//    sum += it.Value()
//}
