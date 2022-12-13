package main

import (
	"bufio"
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
	cranes [stacks]*stack.Stack
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

func ParseStack(scanner *bufio.Scanner) {
	const sep int = 3
	//each line
	for i := 0; i < int(stackheigth); i++ {
		scanner.Scan()
		line := scanner.Text()
		//each columns
		for j := 0; j < int(stacks); j++ {
			c := string(line[1+sep*j+j])
			if c != " " {
				cranes[j].Push(c)
			}
		}
	}
	flipStacks()
	scanner.Scan() //number lines
	scanner.Scan() //empty lines
}

func main() {
	fmt.Printf("Day 5\n")
	initStacks()

	file, err := os.Open("input.txt")
	if err != nil {
		panic("Error: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ParseStack(scanner)

	// foreach command
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue //ignore empty lines
		}
		c := ParseCommand(scanner.Text())
		for i := 0; i < c.qty; i++ {
			cranes[c.to-1].Push(cranes[c.from-1].Pop())
		}
	}
	PrintTopItemOnCranes()
}

func PrintTopItemOnCranes() {
	for i := 0; i < int(stacks); i++ {
		fmt.Print(cranes[i].Pop())
	}
	fmt.Print("\n")
}

func initStacks() {
	for i := 0; i < int(stacks); i++ {
		cranes[i] = stack.New()
	}
}

func flipStacks() {
	for i := 0; i < int(stacks); i++ {
		cranes[i].Flip()
	}
}
