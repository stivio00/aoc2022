package main

import (
	"fmt"
)

func main() {
	grid := Grid{}
	grid.LoadGrid("sample.txt")
	grid.SolveVisibility()
	result1 := grid.CountVisible()
	fmt.Println("Part 1:", result1)
}
