package main

import (
	"fmt"
)

func main() {
	grid := Grid{}
	grid.LoadGrid("input.txt")
	grid.SolveVisibility()
	result1 := grid.CountVisible()
	fmt.Println("Part 1:", result1)
}
