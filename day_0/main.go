package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"bubunyo/aoc-25/common"
)

const prefix = "day_3"

func main() {
	prefix := "Day 1:"
	fmt.Println(prefix, "Demo 1-", run_("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run_("input"))
	fmt.Println(prefix, "Demo 2 ->", run_("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run_("input"))
}

func run_(fp string) any {
	acc := 0
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		s := strings.Split(str, " ")
		i, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatalln("li", err)
		}
		_ = i
	}
	return acc
}
