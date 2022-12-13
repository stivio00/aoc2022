package main

import (
	"reflect"
	"testing"
)

func TestCircularBuffer(t *testing.T) {
	cb := *NewCB(4)
	Assert(t, 12, 12)

	Assert(t, cb._data, []byte{0, 0, 0, 0})

	cb.Append(42)
	Assert(t, cb._data, []byte{42, 0, 0, 0})

	cb.Append(7)
	Assert(t, cb._data, []byte{42, 7, 0, 0})

	cb.Append(0)
	Assert(t, cb._data, []byte{42, 7, 0, 0})

	cb.Append(42)
	Assert(t, cb._data, []byte{42, 7, 0, 42})

	cb.Append(5)
	Assert(t, cb._data, []byte{5, 7, 0, 42})

}

func TestCechDiff(t *testing.T) {
	cb := *NewCB(4)
	cb.Append('s')
	cb.Append('s')
	cb.Append('p')
	cb.Append('r')
	Assert(t, cb.CheckIfAllAreDiff(), false)
}

func TestCechDiff2(t *testing.T) {
	cb := *NewCB(4)
	cb.Append('s')
	cb.Append('x')
	cb.Append('p')
	cb.Append('r')
	Assert(t, cb.CheckIfAllAreDiff(), true)
}

// Util
func Assert(t *testing.T, actual interface{}, expected interface{}) {
	t1 := reflect.TypeOf(actual)
	t2 := reflect.TypeOf(expected)
	if t1 != t2 {
		t.Fatalf("Types are diferent: %T != %T", actual, expected)
	}
	if actual != expected {
		t.Fatalf("Values are diferent: %v != %v", actual, expected)
	}
}

func TestPermutations(t *testing.T) {
	x := []int{0, 1, 2, 3}

	out := make([][2]int, 0)
	for i, v := range x {
		for j := i + 1; j < len(x); j++ {
			out = append(out, [2]int{v, x[j]})
		}
	}
	t.Log(x)
}
