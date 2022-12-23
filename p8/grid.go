package main

import (
	"bufio"
	"os"
)

type Direction uint8

const (
	Left Direction = iota
	Rigth
	Top
	Down
)

type State uint8

const (
	Visible State = iota
	NotVisible
	Unknow
)

type Grid struct {
	Data   [][]int
	Vis    [][]bool
	Width  int
	Heigth int
}

// [x][y] : x=vertical y=horizontal , 0,0=topleftcorner
func (g *Grid) LoadGrid(filename string) {
	g.fillGrid(filename)
	g.recalculateDimensions()
	g.createEmptyVisibilityGrid()
}

func (g *Grid) fillGrid(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	g.Data = make([][]int, 0)
	scanner := bufio.NewScanner(file)

	var x int = 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		g.Data = append(g.Data, make([]int, 0))
		for _, t := range scanner.Text() {
			g.Data[x] = append(g.Data[x], int(t-'0'))
		}
		x += 1
	}
}

func (g *Grid) recalculateDimensions() {
	g.Heigth = len(g.Data)
	g.Width = len(g.Data[0])
}

func (g *Grid) createEmptyVisibilityGrid() {
	g.Vis = make([][]bool, g.Heigth)
	for i := 0; i < g.Heigth; i++ {
		g.Vis[i] = make([]bool, g.Width)
	}
}

// Part 1, only check around
func (g *Grid) getVisibility(x, y int, direction Direction) bool {
	if g.isEdge(x, y) {
		return true
	}
	currrentTree, _ := g.current(x, y)

	if direction == Top {
		val, vis := g.current(x-1, y)
		return currrentTree >= val && vis
	}
	if direction == Down {
		val, vis := g.current(x+1, y)
		return currrentTree >= val && vis
	}
	if direction == Left {
		val, vis := g.current(x, y-1)
		return currrentTree >= val && vis
	}
	if direction == Rigth {
		val, vis := g.current(x, y+1)
		return currrentTree >= val && vis
	}

	return false
}

func (g *Grid) isEdge(x, y int) bool {
	if x == 0 || x == g.Heigth-1 {
		return true
	}
	if y == 0 || y == g.Width-1 {
		return true
	}
	return false
}

func (g *Grid) current(x, y int) (int, bool) {
	return g.Data[x][y], g.Vis[x][y]
}

// iterate ech rectangle from outer to inside.
func (g *Grid) SolveVisibility() {
	for i := 0; i < g.Heigth/2-1; i++ {
		g.rectangleVisibility(i, i)
	}
}

func (g *Grid) rectangleVisibility(x, y int) {
	i := x
	//Top line
	for j := y; j < g.Width-y; j++ {
		g.Vis[i][j] = g.getVisibility(i, j, Top)
	}
	//Left line
	j := y
	for i := x; i < g.Heigth-x; i++ {
		g.Vis[i][j] = g.getVisibility(i, j, Left)
	}
	//Rigth line
	j = g.Width - 1 - x
	for i := x; i < g.Heigth-y; i++ {
		g.Vis[i][j] = g.getVisibility(i, j, Rigth)
	}
	//Down rigth
	i = g.Heigth - x - 1
	for j := y; j < g.Width-y; j++ {
		g.Vis[i][j] = g.getVisibility(i, j, Down)
	}
}

func (g *Grid) CountVisible() int {
	sum := 0
	for i := 0; i < g.Heigth; i++ {
		for j := 0; j < g.Width; j++ {
			if _, vis := g.current(i, j); vis {
				sum += 1
			}
		}
	}
	return sum
}
