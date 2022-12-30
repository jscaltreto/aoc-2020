package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var values []int
	loadInput(&values)
	a(values)
	b(values)
}

func loadInput(values *[]int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		integer, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		*values = append(*values, integer)
	}

}

func a(values []int) {
	for i, v := range values {
		for _, v2 := range values[i+1 : len(values)] {
			if v+v2 == 2020 {
				fmt.Printf("Found Values making 2020: %d %d\n", v, v2)
				fmt.Printf("Product: %d\n", v*v2)
				return
			}
		}
	}
}

func b(values []int) {
	for i, v := range values {
		for _, v2 := range values[i+1 : len(values)] {
			if v+v2 < 2020 {
				for _, v3 := range values {
					if v+v2+v3 == 2020 {
						fmt.Printf("Found Values making 2020: %d %d %d\n", v, v2, v3)
						fmt.Printf("Product: %d\n", v*v2*v3)
						return
					}
				}
			}
		}
	}
}
