package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	calories []int = make([]int, 1)
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentElve := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currentElve++
			calories = append(calories, 0)
			continue
		}
		calorie, _ := strconv.Atoi(line)
		calories[currentElve] += calorie
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j] //desc sort
	})

	fmt.Println("Part1: ", calories[0])
	fmt.Println("Part2: ", Sum(calories[0:3]))
}

func Sum(slice []int) (sum int) {
	for _, v := range slice {
		sum += v
	}
	return
}
