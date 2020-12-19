package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const (
	INPUT = "input.txt"
)

var curTime int
var busNumbers []int

var bBusNumbers []*big.Int
var bRemainders []*big.Int

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
	curTime, _ = strconv.Atoi(strings.Split(string(data), "\n")[0])
	for o, b := range strings.Split(strings.Split(string(data), "\n")[1], ",") {
		busNum, _ := strconv.Atoi(b)
		busNumbers = append(busNumbers, busNum)
		if b != "x" {
			bBusNumbers = append(bBusNumbers, big.NewInt(int64(busNum)))
			bRemainders = append(bRemainders, big.NewInt(int64(busNum-o)))
		}
	}
}

func a() {
	minutesToWait, busNumber := curTime, 0
	for _, b := range busNumbers {
		if b == 0 {
			continue
		}
		_, frac := math.Modf(float64(curTime) / float64(b))
		minutesLeft := b - int(float64(b)*frac)
		if minutesLeft < minutesToWait {
			minutesToWait = minutesLeft
			busNumber = b
		}
	}
	fmt.Println("Part A:", (minutesToWait-1)*busNumber)
}

func b() {
	bigN := big.NewInt(1)
	for _, n1 := range bBusNumbers {
		bigN.Mul(bigN, n1)
	}

	var x, bigN1, s, z big.Int
	for i, n1 := range bBusNumbers {
		bigN1.Div(bigN, n1)
		z.GCD(nil, &s, n1, &bigN1)
		x.Add(&x, s.Mul(bRemainders[i], s.Mul(&s, &bigN1)))
	}
	fmt.Println("Part B:", x.Mod(&x, bigN))
}
