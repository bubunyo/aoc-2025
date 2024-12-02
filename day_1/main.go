package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"bubunyo/aoc-25/common"
)

func main() {
	prefix := "Day 1:"
	fmt.Println(prefix, "Demo 1-", run1("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run1("input"))
	fmt.Println(prefix, "Demo 2 ->", run2("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run2("input"))
}

func run2(fp string) any {
	left, right := []int{}, map[int]int{}
	for str := range common.IterateFileContent("day_1/" + fp) {
		s := strings.Split(str, "   ")
		li, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatalln("li", err)
		}
		left = append(left, li)

		ri, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatalln("ri", err)
		}
		if _, ok := right[ri]; ok {
			right[ri] += 1
		} else {
			right[ri] = 1
		}
	}
	acc := 0
	for i := 0; i < len(left); i++ {
		l := left[i]
		r := right[l]
		acc += (l * r)
	}
	return acc
}

func run1(fp string) any {
	left, right := []int{}, []int{}
	for str := range common.IterateFileContent("day_1/" + fp) {
		s := strings.Split(str, "   ")
		li, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatalln("li", err)
		}
		left = append(left, li)

		ri, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatalln("ri", err)
		}
		right = append(right, ri)
	}
	sort.Ints(left)
	sort.Ints(right)
	acc := 0
	for i := 0; i < len(left); i++ {
		v := left[i] - right[i]
		if v <= 0 {
			v = v * -1
		}
		acc += v
	}
	return acc
}
