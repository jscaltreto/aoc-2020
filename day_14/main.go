package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	INPUT = "input.txt"
	ONE   = '1'
	ZERO  = '0'
	X     = 'X'
)

type Memory map[uint64]uint64

var instructions []string

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
	instructions = strings.Split(string(data), "\n")
}

func readInst(line string) (address uint64, value uint64) {
	instRe, _ := regexp.Compile(`^mem\[(\d+)\] = (\d+)$`)
	match := instRe.FindAllStringSubmatch(line, -1)
	if match == nil {
		log.Fatalf("Invalid Instruction: %s", line)
	}
	addrInt, _ := strconv.Atoi(match[0][1])
	valInt, _ := strconv.Atoi(match[0][2])
	address, value = uint64(addrInt), uint64(valInt)
	return
}

func readMaskA(line string) (maskOne uint64, maskZero uint64) {
	maskOne, maskZero = 0, 0
	for _, bit := range line[7:] {
		maskOne <<= 1
		maskZero <<= 1
		switch bit {
		case ONE:
			maskOne |= 1
		case ZERO:
			maskZero |= 1
		}
	}
	return
}

func a() {
	var memory Memory = make(Memory)
	var maskOne, maskZero uint64
	for _, line := range instructions {
		switch line[:3] {
		case "mas":
			maskOne, maskZero = readMaskA(line)
		case "mem":
			address, value := readInst(line)
			newValue := (^value | ^maskZero) & (value | maskOne)
			memory[address] = newValue
		}
	}
	total := uint64(0)
	for _, value := range memory {
		total += value
	}
	fmt.Println("Part A:", total)
}

func readMaskB(line string) (maskOne uint64, floating uint64, fbMasks map[uint64]bool) {
	maskOne, floating = 0, 0
	checkNums := []uint64{0}
	fbMasks = map[uint64]bool{0: true}
	for idx, bit := range line[7:] {
		maskOne <<= 1
		floating <<= 1
		switch bit {
		case ONE:
			maskOne |= 1
		case X:
			floating |= 1
			checkNums = append(checkNums, uint64(1<<(35-idx)))
		}
	}
	for {
		checkNum := checkNums[0]
		checkNums = checkNums[1:]
		for _, combine := range checkNums {
			newNum := checkNum | combine
			if _, ok := fbMasks[newNum]; !ok {
				checkNums = append(checkNums, newNum)
				fbMasks[newNum] = true
			}
		}
		if len(checkNums) == 0 {
			break
		}

	}
	return
}

func b() {
	var memory Memory = make(Memory)
	var baseMask, floatingBits uint64
	var fbMasks map[uint64]bool
	for _, line := range instructions {
		switch line[:3] {
		case "mas":
			baseMask, floatingBits, fbMasks = readMaskB(line)
		case "mem":
			baseAddress, value := readInst(line)
			baseAddress |= baseMask
			for maskOne := range fbMasks {
				maskZero := ^maskOne & floatingBits
				newAddress := (^baseAddress | ^maskZero) & (baseAddress | maskOne)
				memory[newAddress] = value
			}
		}
	}
	total := uint64(0)
	for _, value := range memory {
		total += value
	}
	fmt.Println("Part B:", total)
}
