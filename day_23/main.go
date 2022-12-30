package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT = "input.txt"

func main() {
	a()
	b()
}

func LoadInput() string {
	data, err := os.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(data), "\n")
}

type Cup struct {
	label int
	next  *Cup
}

func addCup(cupNum int, lastCup *Cup) *Cup {
	cup := Cup{cupNum, nil}
	lastCup.next = &cup
	return &cup
}

func move(curCup *Cup, cupMap map[int]*Cup) *Cup {
	removed := curCup.next
	curCup.next = removed.next.next.next

	destCupLabel := curCup.label - 1
	for destCupLabel == removed.label ||
		destCupLabel == removed.next.label ||
		destCupLabel == removed.next.next.label ||
		destCupLabel == 0 {
		if destCupLabel == 0 {
			destCupLabel = len(cupMap)
		} else {
			destCupLabel--
		}
	}
	destCup := cupMap[destCupLabel]
	removed.next.next.next = destCup.next
	destCup.next = removed

	return curCup.next
}

func a() {
	cupMap := make(map[int]*Cup)

	startCups := strings.Split(LoadInput(), "")
	firstCup, _ := strconv.Atoi(startCups[0])
	curCup := &Cup{firstCup, nil}
	lastCup := curCup
	cupMap[firstCup] = curCup
	for _, cup := range startCups[1:] {
		cupNum, _ := strconv.Atoi(cup)
		lastCup = addCup(cupNum, lastCup)
		cupMap[cupNum] = lastCup
	}
	lastCup.next = curCup

	for i := 0; i < 100; i++ {
		curCup = move(curCup, cupMap)
	}

	newOrder := ""
	curCup = cupMap[1].next
	for curCup != cupMap[1] {
		newOrder += strconv.Itoa(curCup.label)
		curCup = curCup.next
	}
	fmt.Println("Part A:", newOrder)
}

func b() {
	cupMap := make(map[int]*Cup)

	startCups := strings.Split(LoadInput(), "")
	firstCup, _ := strconv.Atoi(startCups[0])
	curCup := &Cup{firstCup, nil}
	lastCup := curCup
	cupMap[firstCup] = curCup
	for _, cup := range startCups[1:] {
		cupNum, _ := strconv.Atoi(cup)
		lastCup = addCup(cupNum, lastCup)
		cupMap[cupNum] = lastCup
	}
	for cupNum := 10; cupNum <= 1000000; cupNum++ {
		lastCup = addCup(cupNum, lastCup)
		cupMap[cupNum] = lastCup
	}
	lastCup.next = curCup

	for i := 0; i < 10000000; i++ {
		curCup = move(curCup, cupMap)
	}

	product := cupMap[1].next.label * cupMap[1].next.next.label
	fmt.Println("Part B:", product)
}
