package main

import (
	"bufio"
	"container/list"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string

type Operation int

const (
	Noop Operation = iota
	AddX
)

type Registers struct {
	X int
}

type Instruction struct {
	opcode     Operation
	arg1       int
	cyclesLeft int
}

type CPU struct {
	registers            Registers
	executingIntructions *list.List
	currentCycle         int
}

type SignalStrengths []int

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	cpu := &CPU{
		registers:            Registers{X: 1},
		executingIntructions: list.New(),
		currentCycle:         0,
	}
	cpu.executingIntructions.Init()
	stregths := make(SignalStrengths, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		instruction := parseLine(line)
		cpu.Fetch(instruction)
		stregths.Inspect(cpu)
		fmt.Println(cpu.currentCycle, cpu.registers.X)
	}

	fmt.Println("Part 1: ", stregths.sum())
}

func parseLine(line string) *Instruction {
	var opcode string
	var arg int
	fmt.Sscanf(line, "%s %d", &opcode, &arg)

	if opcode == "addx" {
		return &Instruction{AddX, arg, 2}
	} else if opcode == "noop" {
		return &Instruction{Noop, 0, 1}
	}
	return nil
}

func (cpu *CPU) Fetch(i *Instruction) {
	cpu.executingIntructions.PushBack(i)
	cpu.execute()
}

func (cpu *CPU) parallelExecute() {

	removal := []*list.Element{}

	for e := cpu.executingIntructions.Front(); e != nil; e = e.Next() {
		i := e.Value.(*Instruction)
		i.cyclesLeft -= 1
		if i.cyclesLeft < 0 {
			cpu.registers.X += i.arg1
			removal = append(removal, e)
		}
	}

	for _, e := range removal {
		cpu.executingIntructions.Remove(e)
	}
	cpu.currentCycle += 1
}

func (cpu *CPU) execute() {

	removal := []*list.Element{}

	for e := cpu.executingIntructions.Front(); e != nil; e = e.Next() {
		i := e.Value.(*Instruction)
		i.cyclesLeft -= 1
		if i.cyclesLeft < 0 {
			cpu.registers.X += i.arg1
			removal = append(removal, e)
		}
		break
	}

	for _, e := range removal {
		cpu.executingIntructions.Remove(e)
	}
}

func (s *SignalStrengths) sum() (result int) {
	for _, v := range *s {
		result += v
	}
	return
}

func (s *SignalStrengths) Inspect(cpu *CPU) {
	if (cpu.currentCycle+20)%40 == 0 {
		*s = append([]int(*s), cpu.registers.X*cpu.currentCycle)
	}
}
