package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	INPUT = "input.txt"
)

type Memory map[uint64]uint64

var startNums []int

func main() {
	loadInput()
	a()
	b()
}

func loadInput() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range strings.Split(string(data), ",") {
		intNum, _ := strconv.Atoi(n)
		startNums = append(startNums, intNum)
	}
}

func checkTurn(maxTurns int) int {
	turnTrack := make(map[int]int)
	nextNum, lastNum := startNums[len(startNums)-1], 0
	for idx, num := range startNums[:len(startNums)-1] {
		turnTrack[num] = idx + 1
	}
	for turnNum := len(startNums); turnNum <= maxTurns; turnNum++ {
		lastNum = nextNum
		if lastTurn, ok := turnTrack[lastNum]; ok {
			nextNum = turnNum - lastTurn
		} else {
			nextNum = 0
		}
		turnTrack[lastNum] = turnNum

	}
	return lastNum
}
func a() {
	lastNum := checkTurn(2020)
	fmt.Println("Part A:", lastNum)
}
func b() {
	lastNum := checkTurn(30000000)
	fmt.Println("Part B:", lastNum)
}
