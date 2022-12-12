// Author Stephen Krol
package main

import (
	"bufio"
	"fmt"
	"log"
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

func ParseLine(line string) *ElvesPair {
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
	var overlapperElves int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elvePair := ParseLine(scanner.Text())
		if elvePair.FullContains() {
			overlapperElves += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result = %#v", overlapperElves)

}
