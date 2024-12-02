package main

import (
	"fmt"

	"bubunyo/aoc-25/common"
)

func main() {
	prefix := "Day 1:"
	fmt.Println(prefix, "Demo 1-", run_("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run_("input"))
	fmt.Println(prefix, "Demo 2 ->", run_("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run_("input"))
}

func run_(fp string) any {
	for str := range common.IterateFileContent("day_0/" + fp) {
		_ = str
	}
	acc := 0
	return acc
}
