package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type password struct {
	min      int
	max      int
	char     string
	password string
}

func main() {
	var passwords []password
	loadInput(&passwords)

	a(passwords)
	b(passwords)
}

func loadInput(values *[]password) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	re, _ := regexp.Compile(`^(\d+)-(\d+) (\w): (\w+)$`)
	for scanner.Scan() {
		match := re.FindAllStringSubmatch(scanner.Text(), -1)
		if match != nil {
			min, _ := strconv.Atoi(match[0][1])
			max, _ := strconv.Atoi(match[0][2])
			*values = append(*values, password{
				min:      min,
				max:      max,
				char:     match[0][3],
				password: match[0][4],
			})
		}
	}
}

func a(values []password) {
	valid := 0
	for _, v := range values {
		count := 0
		for _, c := range v.password {
			if string(c) == v.char {
				count++
			}
		}
		if v.min <= count && v.max >= count {
			valid++
		}
	}
	fmt.Printf("Valid passwords Part 1: %d\n", valid)
}

func b(values []password) {
	valid := 0
	for _, v := range values {
		chars := []rune(v.password)

		if len(v.password) >= v.max {
			matches := 0
			if string(chars[v.min-1]) == v.char {
				matches++
			}
			if string(chars[v.max-1]) == v.char {
				matches++
			}
			if matches == 1 {
				valid++
			}
		}
	}
	fmt.Printf("Valid passwords Part 2: %d\n", valid)
}
