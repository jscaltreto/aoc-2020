package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string
type passports []passport

func main() {
	var p passports
	loadInput(&p)
	a(p)
	b(p)
}

func flush(pass *passport, p *passports) {
	*p = append(*p, *pass)
	*pass = make(passport)

}

func loadInput(p *passports) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	pass := make(passport)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			flush(&pass, p)
			continue
		}
		for _, f := range strings.Fields(line) {
			kv := strings.SplitN(f, ":", 2)
			pass[kv[0]] = kv[1]
		}
	}
	flush(&pass, p)
}

func a(p passports) {
	required := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	valid := 0
PASS:
	for _, pass := range p {
		isValid := true
		for _, key := range required {
			if _, ok := pass[key]; !ok {
				continue PASS
			}
		}
		if isValid {
			valid++
		}

	}
	fmt.Printf("Valid: %d\n", valid)
}

func checkYear(s string, min int, max int) bool {
	re, _ := regexp.Compile(`^(\d{4})$`)
	match := re.FindAllStringSubmatch(s, -1)
	if match != nil {
		year, _ := strconv.Atoi(match[0][1])
		if year >= min && year <= max {
			return true
		}
	}
	return false
}

func byr(s string) bool {
	return checkYear(s, 1920, 2002)
}

func iyr(s string) bool {
	return checkYear(s, 2010, 2020)
}

func eyr(s string) bool {
	return checkYear(s, 2020, 2030)
}

func hgt(s string) bool {
	re, _ := regexp.Compile(`^(\d+)(cm|in)$`)
	match := re.FindAllStringSubmatch(s, -1)
	if match != nil {
		hgt, _ := strconv.Atoi(match[0][1])
		if match[0][2] == "cm" && hgt >= 150 && hgt <= 193 {
			return true
		}
		if match[0][2] == "in" && hgt >= 59 && hgt <= 76 {
			return true
		}
	}
	return false
}

func hcl(s string) bool {
	re, _ := regexp.Compile(`^#[0-9a-f]{6}$`)
	match := re.FindAllStringSubmatch(s, -1)
	if match != nil {
		return true
	}
	return false
}

func ecl(s string) bool {
	re, _ := regexp.Compile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	match := re.FindAllStringSubmatch(s, -1)
	if match != nil {
		return true
	}
	return false
}

func pid(s string) bool {
	re, _ := regexp.Compile(`^\d{9}$`)
	match := re.FindAllStringSubmatch(s, -1)
	if match != nil {
		return true
	}
	return false
}

func b(p passports) {
	required := map[string]func(string) bool{
		"byr": byr,
		"iyr": iyr,
		"eyr": eyr,
		"hgt": hgt,
		"hcl": hcl,
		"ecl": ecl,
		"pid": pid,
	}
	valid := 0
PASS:
	for _, pass := range p {
		for key, fn := range required {
			if _, ok := pass[key]; !ok {
				continue PASS
			}
			if !fn(pass[key]) {
				continue PASS
			}
		}
		valid++
	}
	fmt.Printf("Valid: %d\n", valid)
}
