package main

import (
	"fmt"
	"strings"

	"bubunyo/aoc-25/common"
)

const prefix = "day_8"

func main() {
	fmt.Println(prefix, "Demo 1-", run1("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run1("input"))
	fmt.Println(prefix, "Demo 2 ->", run2("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run2("input"))
}

func run2(fp string) any {
	w, h := 0, 0
	n := map[string][][2]int{}
	m := [][]string{}
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		s := strings.Split(str, "")
		m = append(m, s)
		w = len(s)
		for i, v := range s {
			if v != "." {
				_, ok := n[v]
				if ok {
					n[v] = append(n[v], [2]int{i, h})
				} else {
					n[v] = [][2]int{{i, h}}
				}
			}
		}
		h++
	}

	validP := func(p [2]int) bool {
		if p[0] < 0 || p[1] < 0 {
			return false
		}
		if p[0] >= w {
			return false
		}
		if p[1] >= h {
			return false
		}
		return true

	}

	an := map[[2]int]struct{}{}
	_ = an

	for _, v := range n {
		for i := 0; i < len(v); i++ {
			for j := i; j < len(v); j++ {
				if i == j {
					continue
				}
				ix := v[i][0]
				iy := v[i][1]
				jx := v[j][0]
				jy := v[j][1]

				diffX := ix - jx
				if ix < jx {
					diffX = jx - ix
				}
				diffY := iy - jy
				if iy < jy {
					diffY = jy - iy
				}

				if jx > ix {
					diffX *= -1
				}

				a1x := ix
				a2x := jx
				a1y := iy
				a2y := jy

				an[v[i]] = struct{}{}
				an[v[j]] = struct{}{}

				for {

					// outwards

					a1x += diffX
					a2x -= diffX

					a1y -= diffY
					a2y += diffY

					a1 := [2]int{a1x, a1y}
					a2 := [2]int{a2x, a2y}

					if !validP(a1) && !validP(a2) {
						break
					}

					if validP(a1) {
						// m[a1[1]][a1[0]] = "#"
						an[a1] = struct{}{}

					}

					if validP(a2) {
						// m[a2[1]][a2[0]] = "#"
						an[a2] = struct{}{}
					}
				}

			}
		}
	}

	return len(an)
}

func run1(fp string) any {
	w, h := 0, 0
	n := map[string][][2]int{}
	m := [][]string{}
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		s := strings.Split(str, "")
		m = append(m, s)
		w = len(s)
		for i, v := range s {
			if v != "." {
				_, ok := n[v]
				if ok {
					n[v] = append(n[v], [2]int{i, h})
				} else {
					n[v] = [][2]int{{i, h}}
				}
			}
		}
		h++
	}

	validP := func(p [2]int) bool {
		if p[0] < 0 || p[1] < 0 {
			return false
		}
		if p[0] >= w {
			return false
		}
		if p[1] >= h {
			return false
		}
		return true
	}

	an := map[[2]int]struct{}{}

	for _, v := range n {
		for i := 0; i < len(v); i++ {
			for j := i; j < len(v); j++ {
				if i == j {
					continue
				}
				ix := v[i][0]
				iy := v[i][1]
				jx := v[j][0]
				jy := v[j][1]

				diffX := ix - jx
				if ix < jx {
					diffX = jx - ix
				}
				diffY := iy - jy
				if iy < jy {
					diffY = jy - iy
				}

				a1x := ix + diffX
				a2x := jx - diffX
				if jx > ix {
					a1x = ix - diffX
					a2x = jx + diffX
				}

				a1y := iy - diffY
				a2y := jy + diffY

				a1 := [2]int{a1x, a1y}
				a2 := [2]int{a2x, a2y}

				if validP(a1) {
					// m[a1[1]][a1[0]] = "#"
					an[a1] = struct{}{}
				}

				if validP(a2) {
					// m[a2[1]][a2[0]] = "#"
					an[a2] = struct{}{}
				}

			}
		}
	}

	return len(an)
}
