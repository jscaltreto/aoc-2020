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

var keys []int

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
	for _, strKey := range strings.Split(string(data), "\n") {
		intKey, _ := strconv.Atoi(strKey)
		keys = append(keys, intKey)
	}
}

func transform(value int, subNum int) int {
	value = value * subNum
	return value % 20201227
}

func a() {
	privateKeys := make([]int, 2)
	for idx, key := range keys {
		loopSize := 0
		value := 1
		for value != key {
			value = transform(value, 7)
			loopSize++
		}
		privateKey := 1
		for i := 0; i < loopSize; i++ {
			privateKey = transform(privateKey, keys[idx^1])
		}
		privateKeys[idx] = privateKey
	}
	if privateKeys[0] != privateKeys[1] {
		log.Fatal("Something went wrong, bro.")
	}
	fmt.Println("Part A:", privateKeys[0])
}

func b() {
	fmt.Println("Part B: Merry XMas!")
}
