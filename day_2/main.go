package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"bubunyo/aoc-25/common"
)

const prefix = "day_2"

func main() {
	// fmt.Println(prefix, "Demo 1::", run1("input_demo"))
	// fmt.Println(prefix, "Run 1 ->", run1("input"))
	fmt.Println(prefix, "Demo 2 ->", run2("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run2("input")) // 663 -> high, 650 - low, 658 -- answer
}

func run2(fp string) any {
	acc := 0
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		_ = str
		report := strings.Split(str, " ")

		prev, err := strconv.Atoi(report[0])
		if err != nil {
			log.Fatalln("strconv:", err)
		}

		level := 0 // 0 for unset,  1 for increasing, -1 for decreasing

		safe := true
		dampCount := 0

		for dampCount <= 1 && safe {

			for i := 1; i < len(report); i++ {
				curr, err := strconv.Atoi(report[i])

				if err != nil {
					log.Fatalln("strconv:", err)
				}

				if curr == prev {
					dampCount++
					if dampCount > 1 {
						safe = false
						break
					}
					i = i - 1
					report = append(report[:i], report[i+1:]...)
					break
				}

				if i == 1 {
					if curr > prev {
						level = 1 // increasing
					} else {
						level = -1 // decreasing
					}
				}

				var amt int
				if level > 0 {
					amt = curr - prev
				} else {
					amt = prev - curr
				}

				if amt <= 0 || amt > 3 {
					dampCount++
					if dampCount > 1 {
						safe = false
						break
					}
					i = i - 1
					report = append(report[:i], report[i+1:]...)
					break
				}
				prev = curr
			}
		}

		if safe {
			acc++
		}
	}
	return acc
}

func run1(fp string) any {
	acc := 0
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		_ = str
		report := strings.Split(str, " ")

		prev, err := strconv.Atoi(report[0])
		if err != nil {
			log.Fatalln("strconv:", err)
		}

		level := 0 // 0 for unset,  1 for increasing, -1 for decreasing

		safe := true

		for i := 1; i < len(report); i++ {
			curr, err := strconv.Atoi(report[i])

			if err != nil {
				log.Fatalln("strconv:", err)
			}

			if curr == prev {
				safe = false
				break
			}

			// set level only on first iteration
			if i == 1 {
				if curr > prev {
					level = 1 // increasing
				} else {
					level = -1 // decreasing
				}
			}

			var amt int
			if level > 0 {
				amt = curr - prev
			} else {
				amt = prev - curr
			}

			if amt <= 0 || amt > 3 {
				safe = false
				break
			}
			prev = curr
		}
		if safe {
			acc++
		}
	}
	return acc
}
