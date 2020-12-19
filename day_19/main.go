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
)

var rules map[int]string
var messages []string

func main() {
	rules = make(map[int]string)
	loadInput()
	a()
	b()
}

func loadInput() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	sections := strings.Split(string(data), "\n\n")

	fieldRe, _ := regexp.Compile(`^(\d+): (("([ab])")|(((\d+)|(\|))( )?)+)$`)
	for _, f := range strings.Split(sections[0], "\n") {
		match := fieldRe.FindAllStringSubmatch(f, -1)
		if match == nil {
			log.Fatalf("Invalid Rule: %s", f)
		}
		ruleNum, _ := strconv.Atoi(match[0][1])
		if match[0][4] != "" {
			rules[ruleNum] = match[0][4]
		} else {
			rules[ruleNum] = match[0][2]
		}
	}

	for _, msg := range strings.Split(sections[1], "\n") {
		messages = append(messages, msg)
	}
}

func parseRule(ruleNum int) string {
	rule := rules[ruleNum]
	if rule == "a" || rule == "b" {
		return rule
	}
	terms := []string{}
	for _, sec := range strings.Split(rule, " | ") {
		secRe := ""
		for _, rs := range strings.Split(sec, " ") {
			rn, _ := strconv.Atoi(rs)
			secRe += parseRule(rn)
		}
		terms = append(terms, "(?:"+secRe+")")
	}
	if len(terms) > 1 {

		return "(?:" + strings.Join(terms, "|") + ")"
	}
	return terms[0]
}

func a() {
	re, err := regexp.Compile("^" + parseRule(0) + "$")
	if err != nil {
		log.Fatal("Regex Invalid!", err)
	}
	matchCount := 0
	for _, msg := range messages {
		m := re.FindString(msg)
		if m != "" {
			matchCount++
		}
	}
	fmt.Println("Part A:", matchCount)
}

func b() {
	rule42 := parseRule(42)
	rule31 := parseRule(31)
	rule0Re := fmt.Sprintf("^(%s+)(%s+)$", rule42, rule31)
	re, err := regexp.Compile(rule0Re)
	if err != nil {
		log.Fatal("Regex Invalid!", err)
	}
	matchCount := 0
	for _, msg := range messages {
		m := re.FindStringSubmatch(msg)
		if m != nil {
			if len(m[1]) > len(m[2]) {
				matchCount++
			}
		}
	}
	fmt.Println("Part B:", matchCount)
}
