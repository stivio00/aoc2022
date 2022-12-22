package main

import (
	"fmt"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Rigth
	Unknow
)

type Rope struct {
	Head        Position
	Tail        Position
	TailHistory History
}

type Command struct {
	Type   Direction
	Number int
}

type Position struct {
	X int
	Y int
}

type History struct {
	Positions []Position
}

func ParseLine(line string) Command {
	var c string
	var n int

	fmt.Sscanf(line, "%s %d", &c, &n)
	switch c {
	case "U":
		return Command{Up, n}
	case "D":
		return Command{Down, n}
	case "L":
		return Command{Left, n}
	case "R":
		return Command{Rigth, n}
	}
	return Command{Unknow, 0}
}

func (rope *Rope) Move(command Command) {
	for i := 0; i < command.Number; i++ {
		rope.moveOne(command.Type)
		rope.recordTailPosition()
	}
}

func (rope *Rope) moveOne(direction Direction) {
	switch direction {
	case Up:
		rope.Head.X += 1
	case Down:
		rope.Head.X -= 1
	case Left:
		rope.Head.Y -= 1
	case Rigth:
		rope.Head.Y += 1
	}

	dx, dy := rope.Head.Distance(rope.Tail)
	dx, dy = normalizeDistance(dx, dy)
	if !rope.Head.isTouching(rope.Tail) {
		rope.Tail.X += dx
		rope.Tail.Y += dy
	}

}

func (rope *Rope) CountTailUniquePositions() int {
	unique := make(map[Position]struct{}, 0)
	for _, position := range rope.TailHistory.Positions {
		unique[position] = struct{}{}
	}
	return len(unique)
}

func (p Position) Distance(o Position) (dx, dy int) {
	dx = p.X - o.X
	dy = p.Y - o.Y
	return
}

func (p Position) isTouching(o Position) bool {
	dx, dy := p.Distance(o)
	if dx > 1 || dx < -1 {
		return false
	}
	if dy > 1 || dy < -1 {
		return false
	}
	return true
}

// clamp values of the distance
func normalizeDistance(dx, dy int) (int, int) {
	if dx > 1 {
		dx = 1
	}
	if dx < -1 {
		dx = -1
	}
	if dy > 1 {
		dy = 1
	}
	if dy < -1 {
		dy = -1
	}
	return dx, dy
}

func (rope *Rope) recordTailPosition() {
	rope.TailHistory.Positions = append(rope.TailHistory.Positions, rope.Tail)
}
