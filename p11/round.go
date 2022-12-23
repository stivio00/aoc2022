package main

import (
	"fmt"
	"strconv"
)

type ThrowInfo struct {
	ItemIndex int
	ToMonkey  int
}

func Round(monkeys *Monkeys) {

	// map[toMonkeyId]-> []Items
	//throws := make(map[int][]int, 0)

	for _, monkey := range *monkeys {
		fmt.Printf("Monkey %d:\n", monkey.Id)
		inspectItems(&monkey)
	}
}

func inspectItems(monkey *Monkey) {

	for itemIndex, item := range monkey.Items {
		fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", item)
		UpdateWorryLevel(monkey, itemIndex)
		MonkeyGetBored(monkey, itemIndex)
		CheckToThrow(monkey, itemIndex)
	}

}

func CheckToThrow(monkey *Monkey, itemIndex int) bool {
	isDivisible := monkey.Items[itemIndex]%monkey.Divisible == 0

	if isDivisible {
		fmt.Printf("    Current worry level is divisible by %d.\n", monkey.Divisible)
		fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", monkey.Items[itemIndex], monkey.IfTrue)
	} else {
		fmt.Printf("    Current worry level is not divisible by %d.\n", monkey.Divisible)
		fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", monkey.Items[itemIndex], monkey.IfFalse)
	}

	return isDivisible
}

func MonkeyGetBored(monkey *Monkey, itemIndex int) {
	monkey.Items[itemIndex] /= 3
	fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", monkey.Items[itemIndex])
}

func UpdateWorryLevel(monkey *Monkey, itemIndex int) {
	if monkey.Operand1 == "old" && monkey.Op == "*" && monkey.Operand2 == "old" {
		monkey.Items[itemIndex] *= monkey.Items[itemIndex]
		fmt.Printf("    Worry level is multiplied by itself to %d.\n", monkey.Items[itemIndex])
	} else if monkey.Operand1 == "old" && monkey.Op == "*" {
		val, _ := strconv.Atoi(monkey.Operand2)
		monkey.Items[itemIndex] *= val
		fmt.Printf("    Worry level is multiplied by %d to %d.\n", val, monkey.Items[itemIndex])
	} else if monkey.Operand1 == "old" && monkey.Op == "+" {
		val, _ := strconv.Atoi(monkey.Operand2)
		monkey.Items[itemIndex] += val
		fmt.Printf("    Worry level increases by %d to %d.\n", val, monkey.Items[itemIndex])
	}
}
