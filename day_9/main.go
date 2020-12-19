package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var cipher []int
	loadInput(&cipher)
	a(cipher, 25)
	b(cipher, 25)
}

func loadInput(cipher *[]int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Invalid Value: %s", line)
		}
		*cipher = append(*cipher, value)
	}
}

func checkValid(num int, lastN []int) bool {
	for _, numA := range lastN {
		checkFor := num - numA
		for _, numB := range lastN {
			if numB == checkFor && numB != numA {
				return true
			}
		}
	}
	return false
}

func checkNumbers(cipher []int, preambleLen int) (int, error) {
	for idx, num := range cipher {
		if idx < preambleLen {
			continue
		}
		if !checkValid(num, cipher[idx-preambleLen:idx]) {
			return num, nil
		}
	}
	return -1, errors.New("All numbers are valid")
}

func a(cipher []int, preambleLen int) {
	num, err := checkNumbers(cipher, preambleLen)
	if err == nil {
		fmt.Printf("First Invalid Number: %d\n", num)
	}
}

func findSequence(cipher []int, num int) []int {
	for idx, sum := range cipher {
		group := make([]int, 1)
		group[0] = sum
		for true {
			idx++
			if idx >= len(cipher) {
				log.Fatal("Group not found!")
			}
			sum += cipher[idx]
			group = append(group, cipher[idx])
			if sum == num {
				return group
			}
			if sum > num {
				break
			}

		}
	}
	return nil
}

func b(cipher []int, preambleLen int) {
	num, _ := checkNumbers(cipher, preambleLen)
	sequence := findSequence(cipher, num)

	min, max := sequence[0], sequence[0]
	for _, val := range sequence[1:] {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Printf("Sum of min & max: %d\n", min+max)

}
