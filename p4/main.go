// AoC 2022 : https://adventofcode.com/2022/day/4
// Author Stephen Krol
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type ElvesPair struct {
	First  Elve
	Second Elve
}

type Elve struct {
	SectionsRange SectionRange
}

type SectionRange struct {
	From int
	To   int
}

func (s *SectionRange) FullyContains(other *SectionRange) bool {
	if s.From <= other.From && s.To >= other.To {
		return true
	}
	return false
}

func (e *ElvesPair) FullContains() bool {
	firstContainsSecond := e.First.SectionsRange.FullyContains(&e.Second.SectionsRange)
	secondContainsFirst := e.Second.SectionsRange.FullyContains(&e.First.SectionsRange)
	return firstContainsSecond || secondContainsFirst
}

func (e *ElvesPair) Overlapped() bool {
	x2 := float64(e.First.SectionsRange.To)
	y1 := float64(e.Second.SectionsRange.From)
	y2 := float64(e.Second.SectionsRange.To)
	x1 := float64(e.First.SectionsRange.From)
	return math.Max(x1, y1) <= math.Min(x2, y2)
}

func ParseLine(line string) *ElvesPair {
	if len(line) < 7 {
		return nil
	}

	elveRangeString1, elveRangeString2, _ := strings.Cut(line, ",")
	elveSectionStringFrom1, elveSectionStringTo1, _ := strings.Cut(elveRangeString1, "-")
	elveSectionStringFrom2, elveSectionStringTo2, _ := strings.Cut(elveRangeString2, "-")

	elveSectionFrom1, _ := strconv.ParseInt(elveSectionStringFrom1, 10, 32)
	elveSectionTo1, _ := strconv.ParseInt(elveSectionStringTo1, 10, 32)

	elveSectionFrom2, _ := strconv.ParseInt(elveSectionStringFrom2, 10, 32)
	elveSectionTo2, _ := strconv.ParseInt(elveSectionStringTo2, 10, 32)

	return &ElvesPair{
		First: Elve{
			SectionsRange: SectionRange{
				From: int(elveSectionFrom1),
				To:   int(elveSectionTo1),
			},
		},
		Second: Elve{
			SectionsRange: SectionRange{
				From: int(elveSectionFrom2),
				To:   int(elveSectionTo2),
			},
		},
	}
}

func main() {
	fmt.Println("Day 4")
	var intersectElves int = 0
	var overlappedElves int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elvePair := ParseLine(scanner.Text())

		if elvePair == nil {
			continue
		}

		if elvePair.FullContains() {
			intersectElves += 1
		}

		if elvePair.Overlapped() {
			overlappedElves += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part 1 result = %#v\n", intersectElves)
	fmt.Printf("part 2 result = %#v\n", overlappedElves)

}
