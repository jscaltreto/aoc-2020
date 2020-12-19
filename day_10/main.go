package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Adapters []int

func main() {
	var adapters Adapters
	loadInput(&adapters)
	a(adapters)
	b(adapters)
}

func loadInput(adapters *Adapters) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	*adapters = append(*adapters, 0)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Invalid Value: %s", line)
		}
		*adapters = append(*adapters, value)
	}
	sort.Ints(*adapters)
	*adapters = append(*adapters, (*adapters)[len(*adapters)-1]+3)
}

func a(adapters Adapters) {
	currentJolts, diff1, diff3 := 0, 0, 0

	for _, adapter := range adapters[1:] {
		change := adapter - currentJolts
		currentJolts = adapter
		if change > 3 || change < 1 {
			log.Fatal("Invalid Conversion!")
		}
		if change == 1 {
			diff1++
		}
		if change == 3 {
			diff3++
		}
	}
	fmt.Printf("Part A: %d\n", diff1*diff3)

}

func getTargets(adapters *Adapters, pos int) Adapters {
	targets := Adapters{}
	endPos := pos + 4
	if endPos >= len(*adapters) {
		endPos = len(*adapters) - 1
	}
	for idx, num := range (*adapters)[pos+1 : endPos] {
		if idx+pos < len(*adapters)-1 && num <= (*adapters)[pos]+3 {
			targets = append(targets, num)
		}
	}
	return targets
}

func b2(adapters Adapters) {
	adapterMap := make(map[int]Adapters)
	adapters = append(Adapters{0}, adapters...)
	for idx, adapter := range adapters[:len(adapters)-1] {
		adapterMap[adapter] = getTargets(&adapters, idx)
	}
	permutations := make(map[int]int)
	permutations[adapters[len(adapters)-2]] = 1

	for idx := len(adapters) - 3; idx >= 0; idx-- {
		adapter := adapters[idx]
		children := adapterMap[adapter]
		p := 0
		for _, child := range children {
			p += permutations[child]
		}
		permutations[adapter] = p
	}

	fmt.Printf("Total Arrangements: %d\n", permutations[0])
}

func b(adapters Adapters) {
	permutations := []int{1}

	for idx := 1; idx < len(adapters); idx++ {
		p := 0
		for i := idx - 3; i < idx; i++ {
			if i < 0 {
				continue
			}
			if adapters[i]+3 >= adapters[idx] {
				p += permutations[i]
			}
		}
		permutations = append(permutations, p)
	}
	fmt.Printf("Total Arrangements: %d\n", permutations[len(permutations)-1])
}
