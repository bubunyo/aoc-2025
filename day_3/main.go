package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"bubunyo/aoc-25/common"
)

const prefix = "day_3"

func main() {
	fmt.Println(prefix, "Demo 1-", run1("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run1("input"))
	fmt.Println(prefix, "Demo 2 ->", run2("input_demo2"))
	fmt.Println(prefix, "Run 2 ->", run2("input2"))
}

func run2(fp string) any {
	acc := 0
	var totalStr string
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		totalStr += str
	}

	delRe := regexp.MustCompile(`don't\(\).*?do\(\)`)

	str := delRe.ReplaceAllString(totalStr, "")

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(str, -1)

	// Print matches
	for _, match := range matches {
		fmt.Printf("Full Match: %s, Num1: %s, Num2: %s\n", match[0], match[1], match[2])

		m1, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalln("strconv:", err)
		}
		m2, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalln("strconv:", err)
		}
		acc += (m1 * m2)
	}
	return acc
}

func run1(fp string) any {
	acc := 0
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		// Find all matches
		matches := re.FindAllStringSubmatch(str, -1)

		// Print matches
		for _, match := range matches {

			m1, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalln("strconv:", err)
			}
			m2, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatalln("strconv:", err)
			}
			acc += (m1 * m2)
		}
	}
	return acc
}
