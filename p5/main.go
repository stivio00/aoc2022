package main

import (
	"fmt"
	"os"

	"github.com/stivio00/aoc2022/p5/stack"
)

const (
	stackheigth uint   = 8
	stacks      uint   = 9
	command     string = "move %d from %d to %d"
)

var (
	cranes [stacks]string
)

type CommandMove struct {
	qty  int
	from int
	to   int
}

func ParseCommand(line string) *CommandMove {
	var _qty, _from, _to int
	n, err := fmt.Sscanf(line, command, &_qty, &_from, &_to)
	if n != 3 {
		msg := fmt.Sprintf("Error: parsed %d instead of 3. Value=%s", n, line)
		panic(msg)
	}
	if err != nil {
		panic(err.Error())
	}
	return &CommandMove{
		qty:  _qty,
		from: _from,
		to:   _to,
	}
}

func ParseStack(f *os.File) {

}

func main() {
	fmt.Printf("Day 5\n")

	file, err := os.Open("input.txt")
	if err != nil {
		panic("Error: " + err.Error())
	}
	defer file.Close()

	ParseStack(file)

	c := ParseCommand("move 1 from 3 to 6")
	fmt.Println(c)

	s := stack.New()
	s.Push("W")
	s.Push(42)
	s.Print()
	s.Pop()
	s.Print()
	s.Clear()
	s.Print()

	s.Push(9)
	s.Push(12)
	s.Print()
	s.Flip()
	s.Print()

}
