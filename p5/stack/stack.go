package stack

import "fmt"

var (
	initialCap = 64
)

type Stack struct {
	data []interface{}
	size uint32
}

func New() *Stack {
	_data := make([]interface{}, initialCap)
	return &Stack{
		data: _data,
		size: 0,
	}
}

func (s *Stack) NewCopy() *Stack {
	var _data []interface{}
	copy(_data, s.data)
	return &Stack{_data, s.size}
}

func (s *Stack) Clear() {
	s.size = 0
}

func (s *Stack) Pop() interface{} {
	s.size -= 1
	return s.data[s.size]
}

func (s *Stack) Push(item interface{}) uint32 {
	//TODO: check slice capacity, or len.  and do append if neccesary
	s.data[s.size] = item
	s.size += 1
	return s.size
}

func (s *Stack) Size() uint32 {
	return s.size
}

// non standard stack method
func (s *Stack) Flip() {
	for i, j := uint32(0), s.size-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}

func (s *Stack) Print() {
	fmt.Print("Stack(size:", s.size, ") ")
	for i := uint32(0); i < s.size; i++ {
		fmt.Print(s.data[i], " ")
	}
	fmt.Print("\n")
}

// PART2 of the puzzle
func (s *Stack) PopN(n uint32) []interface{} {
	//TODO POP in the same order
	return nil
}

func (s *Stack) PushN([]interface{}) {
	//TODO Push in the same order

}
