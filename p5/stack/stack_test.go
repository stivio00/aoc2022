package stack

import (
	"testing"
)

func TestStackCreation(t *testing.T) {
	s := New()
	t.Log("lol", s)
	if s.Size() != 0 {
		t.Error("Size should bew size 0")
	}
}

func TestStackPush(t *testing.T) {
	s := New()
	s.Push(41)
	s.Push(42)

	if s.Size() != 2 {
		t.Error("size mismatch")
	}

	if s.Pop() != 42 {
		t.Error()
	}

	if s.Pop() != 41 {
		t.Error()
	}
}

func TestStackPrint(t *testing.T) {
	s := New()
	s.Push(41)
	s.Push(42)
	got := s.String()
	want := "Stack(size:2) 41 42 "
	t.Log(got)
	if got != want {
		t.Error()
	}
}
