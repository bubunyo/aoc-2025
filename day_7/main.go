package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"bubunyo/aoc-25/common"
)

const prefix = "day_7"

func main() {
	fmt.Println(prefix, "Demo 1-", run1("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run1("input"))
	fmt.Println(prefix, "Demo 2 ->", run2("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run2("input"))
}

func join(a, b int) int {
	place := 1
	for i := b; i > 0; i /= 10 {
		place *= 10
	}
	return a*place + b
}

func combine2(l []int) []int {
	if len(l) == 1 {
		return []int{l[0]}
	}
	res := []int{}
	li := len(l) - 1
	for _, v := range combine2(l[:li]) {
		c1 := v + l[li]
		c2 := v * l[li]
		c3 := join(v, l[li])
		res = append(res, c1, c2, c3)
	}
	return res
}

func run2(fp string) any {
	acc := 0
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		s := strings.Split(str, " ")
		tot, err := strconv.Atoi(s[0][:len(s[0])-1])
		if err != nil {
			log.Fatalln(err)
		}
		ins := make([]int, len(s)-1)
		for i := 1; i < len(s); i++ {
			v, err := strconv.Atoi(s[i])
			if err != nil {
				log.Fatalln(err)
			}
			ins[i-1] = v
		}

		for _, v := range combine2(ins) {
			if v == tot {
				acc += tot
				break
			}
		}
	}
	return acc
}

func combine(l []int) []int {
	if len(l) == 1 {
		return []int{l[0]}
	}
	res := []int{}
	li := len(l) - 1
	for _, v := range combine(l[:li]) {
		res = append(res, v+l[li], v*l[li])
	}
	return res
}

func run1(fp string) any {
	acc := 0
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		s := strings.Split(str, " ")
		tot, err := strconv.Atoi(s[0][:len(s[0])-1])
		if err != nil {
			log.Fatalln("li", err)
		}
		ins := make([]int, len(s)-1)
		for i := 1; i < len(s); i++ {
			v, err := strconv.Atoi(s[i])
			if err != nil {
				log.Fatalln("li", err)
			}
			ins[i-1] = v
		}

		for _, v := range combine(ins) {
			if v == tot {
				acc += tot
				break
			}
		}
	}
	return acc
}
