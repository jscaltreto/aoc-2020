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
	PLUS  = iota
	TIMES = iota
)

var expressions []string

type Solver func(string) string

func eval(e string, solver Solver) (string, int) {
	myExp := ""
	i := 0
	for i = 0; i < len(e); i++ {
		c := e[i]
		if c == '(' {
			sec, chars := eval(e[i+1:], solver)
			myExp += sec
			i += chars
		} else if c == ')' {
			i++
			break
		} else {
			myExp += string(c)
		}
	}
	return solver(myExp), i
}

func solveA(myExp string) string {
	cmp := strings.Split(myExp, " ")
	ans, _ := strconv.Atoi(cmp[0])
	var oper int
	for _, item := range cmp[1:] {
		if item == "+" {
			oper = PLUS
		} else if item == "*" {
			oper = TIMES
		} else {
			n, _ := strconv.Atoi(item)
			if oper == PLUS {
				ans += n
			} else if oper == TIMES {
				ans *= n
			}

		}
	}
	return strconv.Itoa(ans)
}

func solveB(myExp string) string {
	cmp := strings.Split(myExp, " ")
	for {
		hasPlus := false
		for term, item := range cmp {
			if item == "+" {
				t1, _ := strconv.Atoi(cmp[term-1])
				t2, _ := strconv.Atoi(cmp[term+1])
				newVal := t1 + t2
				newCmp := []string{}
				if term > 1 {
					newCmp = append(newCmp, cmp[:term-1]...)
				}
				newCmp = append(newCmp, strconv.Itoa(newVal))
				newCmp = append(newCmp, cmp[term+2:]...)
				cmp = newCmp
				hasPlus = true
				break
			}
		}
		if !hasPlus {
			break
		}
	}
	ans, _ := strconv.Atoi(cmp[0])
	for _, item := range cmp[1:] {
		if item != "*" {
			n, _ := strconv.Atoi(item)
			ans *= n
		}
	}
	return strconv.Itoa(ans)
}

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
	expressions = strings.Split(string(data), "\n")
}

func a() {
	sum := 0
	for _, e := range expressions {
		result, _ := eval(e, solveA)
		intResult, _ := strconv.Atoi(result)
		sum += intResult
	}
	fmt.Println("Part A:", sum)
}

func b() {
	sum := 0
	for _, e := range expressions {
		result, _ := eval(e, solveB)
		intResult, _ := strconv.Atoi(result)
		sum += intResult
	}
	fmt.Println("Part B:", sum)
}
