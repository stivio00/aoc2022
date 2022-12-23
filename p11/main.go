package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

const (
	ROUNDS = 10
)

func main() {
	monkeys := new(Monkeys)
	monkeys.Init(input)
	for i := 0; i < ROUNDS; i++ {
		fmt.Printf("=====  Round %d  ======\n", i)
		Round(monkeys)
	}

}
