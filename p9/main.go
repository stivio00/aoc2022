package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	Part1()
	Part2()
}

func Part1() {
	rope := NewRope(1)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		command := ParseLine(line)
		rope.Move(command)
	}

	fmt.Println("part 1: ", rope.CountTailUniquePositions())
}

func Part2() {
	rope := NewRope(9)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		command := ParseLine(line)
		rope.Move(command)
	}

	fmt.Println("part 2: ", rope.CountTailUniquePositions())
}
