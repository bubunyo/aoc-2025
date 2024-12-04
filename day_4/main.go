package main

import (
	"fmt"
	"strings"

	"bubunyo/aoc-25/common"
)

const prefix = "day_4"

func main() {
	fmt.Println(prefix, "Demo 1:", run1("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run1("input"))
	fmt.Println(prefix, "Demo 2 ->", run2("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run2("input"))
}

func run1(fp string) any {
	m := [][]string{}
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		m = append(m, strings.Split(str, ""))
	}
	return search(m)
}

func run2(fp string) any {
	m := [][]string{}
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		m = append(m, strings.Split(str, ""))
	}
	return search2(m)
}

var dirFunc2 = map[dir]func(x, y int) (int, int){
	ne: func(x, y int) (int, int) { return x + 1, y - 1 },
	e:  func(x, y int) (int, int) { return x + 1, y },
	se: func(x, y int) (int, int) { return x + 1, y + 1 },
	s:  func(x, y int) (int, int) { return x, y + 1 },
	sw: func(x, y int) (int, int) { return x - 1, y + 1 },
	w:  func(x, y int) (int, int) { return x - 1, y },
	nw: func(x, y int) (int, int) { return x - 1, y - 1 },
}

func createMatrix(x, y int) [][]string {
	matrix := make([][]string, x)
	for i := range matrix {
		matrix[i] = make([]string, y)
	}
	return matrix
}

func search2(m [][]string) int {
	acc := 0
	dirVal := func(d dir, x, y int) string {
		xx, yy := dirFunc[d](x, y)
		return m[yy][xx]
	}

	check := func(x, y int) {
		ne_v := dirVal(ne, x, y)
		nw_v := dirVal(nw, x, y)
		sw_v := dirVal(sw, x, y)
		se_v := dirVal(se, x, y)

		if nw_v == "M" && ne_v == "M" &&
			se_v == "S" && sw_v == "S" {
			acc++
			return
		}
		if nw_v == "M" && sw_v == "M" &&
			ne_v == "S" && se_v == "S" {
			acc++
			return
		}
		if se_v == "M" && ne_v == "M" &&
			sw_v == "S" && nw_v == "S" {
			acc++
			return
		}
		if se_v == "M" && sw_v == "M" &&
			ne_v == "S" && nw_v == "S" {
			acc++
			return
		}
	}

	for y := 1; y < len(m)-1; y++ {
		for x := 1; x < len(m[y])-1; x++ {
			if m[y][x] == "A" {
				check(x, y)
			}
		}
	}

	return acc
}

type dir int

const (
	n  dir = iota // 0
	ne            // 1
	e             // 2
	se            // 3
	s             // 4
	sw            // 5
	w             // 6
	nw            // 7
)

var dirFunc = map[dir]func(x, y int) (int, int){
	n:  func(x, y int) (int, int) { return x, y - 1 },
	ne: func(x, y int) (int, int) { return x + 1, y - 1 },
	e:  func(x, y int) (int, int) { return x + 1, y },
	se: func(x, y int) (int, int) { return x + 1, y + 1 },
	s:  func(x, y int) (int, int) { return x, y + 1 },
	sw: func(x, y int) (int, int) { return x - 1, y + 1 },
	w:  func(x, y int) (int, int) { return x - 1, y },
	nw: func(x, y int) (int, int) { return x - 1, y - 1 },
}

func search(m [][]string) int {
	acc := 0
	word := []string{"X", "M", "A", "S"}
	var dfs func(dir dir, i, x, y int)
	dfs = func(d dir, i, x, y int) {
		if y < 0 || y >= len(m) {
			return
		}
		if x < 0 || x >= len(m[y]) {
			return
		}
		if word[i] != m[x][y] {
			return
		}
		if i == 3 {
			acc++
			return
		}
		fn := dirFunc[d]
		x, y = fn(x, y)
		dfs(d, i+1, x, y)
	}
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			if y >= len(m) || x >= len(m[y]) {
				continue
			}
			if m[x][y] == word[0] {
				for d := range dirFunc {
					dfs(d, 0, x, y)
				}
			}
		}
	}
	return acc
}
