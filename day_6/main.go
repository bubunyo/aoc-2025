package main

import (
	"fmt"
	"log"
	"strings"

	"bubunyo/aoc-25/common"
)

const prefix = "day_6"

func main() {
	fmt.Println(prefix, "Demo 1-", run1("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run1("input"))
	fmt.Println(prefix, "Demo 2 ->", run2("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run2("input"))
}

func nextDir(d int, c [2]int) [2]int {
	switch d {
	case up:
		return [2]int{c[0], c[1] - 1}
	case right:
		return [2]int{c[0] + 1, c[1]}
	case down:
		return [2]int{c[0], c[1] + 1}
	case left:
		return [2]int{c[0] - 1, c[1]}
	default:
		log.Fatal("invlid direction")
	}
	return [2]int{-1, -1}
}

func run2(fp string) any {
	m := [][]string{}
	cur := [2]int{}
	y := 0
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		line := strings.Split(str, "")
		m = append(m, line)
		for x, v := range line {
			if v == "^" {
				cur = [2]int{x, y}
			}
		}
		y++
	}

	dir := up
	path := [][2]int{}

	for {
		next := nextDir(dir, cur)

		validX := next[0] >= 0 && next[0] < len(m[0])
		validY := next[1] >= 0 && next[1] < len(m)

		if !validX || !validY {
			break
		}
		// check obstacle
		if m[next[1]][next[0]] == "#" {
			dir = (dir + 1) % 4
			continue
		}

		path = append(path, [2]int{next[0], next[1]})
		cur = next
	}

	acc := 0

	move := func(c int, in [2]int, dir int) ([2]int, int) {
		for i := 0; i < c; {
			n := nextDir(dir, in)
			validX := n[0] >= 0 && n[0] < len(m[0])
			validY := n[1] >= 0 && n[1] < len(m)

			if !validX || !validY {
				return [2]int{}, -1
			}

			if m[n[1]][n[0]] == "#" {
				dir = (dir + 1) % 4
				continue
			}
			in = n
			i++
		}
		return in, dir
	}

	og := path[0]

	visited := map[[2]int]struct{}{}

	for i := 1; i < len(path); i++ {

		pos := path[i]
		m[pos[1]][pos[0]] = "#"

		tor := [2]int{og[0], og[1]}
		har := [2]int{og[0], og[1]}

		tdir := up
		hdir := up

		for {

			tor, tdir = move(1, tor, tdir)
			har, hdir = move(2, har, hdir)

			if tdir == -1 || hdir == -1 {
				break
			}

			if tor == har && tdir == hdir {
				_, ok := visited[pos]
				if !ok {
					acc++
					visited[pos] = struct{}{}

					m[pos[1]][pos[0]] = fmt.Sprintf("%d", acc)

				}
				break
			}
		}

		m[pos[1]][pos[0]] = "."
	}

	// exit
	return acc

}

func run1(fp string) any {
	visited := map[[2]int]struct{}{}
	m := [][]string{}
	cur := [2]int{}
	y := 0
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		line := strings.Split(str, "")
		m = append(m, line)
		for x, v := range line {
			if v == "^" {
				cur = [2]int{x, y}
			}
		}
		y++
	}
	visited[cur] = struct{}{}
	acc := 1
	dir := up
	for {
		var next [2]int
		switch dir {
		case up:
			next = [2]int{cur[0], cur[1] - 1}
		case right:
			next = [2]int{cur[0] + 1, cur[1]}
		case down:
			next = [2]int{cur[0], cur[1] + 1}
		case left:
			next = [2]int{cur[0] - 1, cur[1]}
		default:
			log.Fatal("invlid direction")
		}
		validX := next[0] >= 0 && next[0] < len(m[0])
		validY := next[1] >= 0 && next[1] < len(m)

		if !validX || !validY {
			// exit
			return acc
		}
		// check obstacle
		if m[next[1]][next[0]] == "#" {
			dir = (dir + 1) % 4
			continue
		}
		// check visited
		_, ok := visited[next]
		if !ok {
			acc++
			visited[next] = struct{}{}
		}
		cur = next
	}
}

const (
	up = iota
	right
	down
	left
)
