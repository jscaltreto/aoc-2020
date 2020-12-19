package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
)

const startAlpha = 97 // rune of "a"

func main() {
	groupA := uint32(0)
	groupB := ^uint32(0)
	sumA := 0
	sumB := 0

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Groups separated by blank line
			sumA += bits.OnesCount32(groupA)
			sumB += bits.OnesCount32(groupB)
			groupA = 0
			groupB = ^uint32(0)
			continue
		}
		person := uint32(0)
		for _, c := range line {
			person |= 1 << (c - startAlpha)
		}
		groupA |= person
		groupB &= person
	}
	// Remember to add the last group!
	sumA += bits.OnesCount32(groupA)
	sumB += bits.OnesCount32(groupB)
	fmt.Printf("Any Answered: %d\n", sumA)
	fmt.Printf("All Answered: %d\n", sumB)
}
