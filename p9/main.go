package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string

func main() {
	rope := Rope{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		command := ParseLine(line)
		rope.Move(command)
	}

	fmt.Println("part 1: ", rope.CountTailUniquePositions())

}
