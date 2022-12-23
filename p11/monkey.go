package main

import (
	"regexp"
	"strconv"
	"strings"
)

// Regex capturing groups
const (
	MONKEY_NUMBER = 1
	ITEMS         = 2
	OPERAND1      = 3
	OP            = 4
	OPERAND2      = 5
	DIVISIBLE     = 6
	IF_TRUE       = 7
	IF_FALSE      = 8
)

type Monkeys []Monkey

type Monkey struct {
	Id        int
	Items     []int
	Operand1  string
	Op        string
	Operand2  string
	Divisible int
	IfTrue    int
	IfFalse   int
}

func (m *Monkeys) Init(input string) {
	*m = Monkeys(make([]Monkey, 0))

	const monkeyRegex = `Monkey (\d)+:
  Starting items: ([\d, ]+)
  Operation: new = (\w+) ([*+]+) (\w+)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)`

	regex := *regexp.MustCompile(monkeyRegex)
	matches := regex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {

		id, _ := strconv.Atoi(match[MONKEY_NUMBER])
		items := (strings.Split(match[ITEMS], ","))
		divisible := must(strconv.Atoi(match[DIVISIBLE]))
		*m = append(*m, Monkey{
			Id:        id,
			Items:     forEach(items, func(s string) int { v, _ := strconv.Atoi(strings.Trim(s, " ")); return v }),
			Operand1:  match[OPERAND1],
			Op:        match[OP],
			Operand2:  match[OPERAND1],
			Divisible: divisible,
			IfTrue:    must(strconv.Atoi(match[IF_TRUE])),
			IfFalse:   must(strconv.Atoi(match[IF_FALSE])),
		})

	}

}

func (m *Monkey) Recive(item int) {
	m.Items = append(m.Items, item)
}

// using Go's controversial generics
func forEach[Tinput any, Toutput any](list []Tinput, f func(Tinput) Toutput) (output []Toutput) {
	for _, v := range list {
		output = append(output, f(v))
	}
	return
}

// generic version of must
func must[T any](t T, e error) T {
	if e != nil {
		panic(e.Error())
	}
	return t
}
