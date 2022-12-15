package main

import (
	"bufio"
	"os"
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
func (g *Grid) getVisibility(x, y int) bool {
	if g.isEdge(x, y) {
		return true
	}

	currrentTree, _ := g.current(x, y)

	if val, _ := g.top(x, y); currrentTree > val {
		return true
	}
	if val, _ := g.down(x, y); currrentTree > val {
		return true
	}
	if val, _ := g.left(x, y); currrentTree > val {
		return true
	}
	if val, _ := g.rigth(x, y); currrentTree > val {
		return true
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

func (g *Grid) top(x, y int) (int, bool) {
	return g.Data[x+1][y], g.Vis[x+1][y]
}

func (g *Grid) down(x, y int) (int, bool) {
	return g.Data[x-1][y], g.Vis[x-1][y]
}

func (g *Grid) left(x, y int) (int, bool) {
	return g.Data[x][y-1], g.Vis[x][y-1]
}

func (g *Grid) rigth(x, y int) (int, bool) {
	return g.Data[x][y+1], g.Vis[x][y+1]
}

func (g *Grid) SolveVisibility() {
	for i := 0; i < g.Heigth; i++ {
		for j := 0; j < g.Width; j++ {
			g.Vis[i][j] = g.getVisibility(i, j)
		}
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
