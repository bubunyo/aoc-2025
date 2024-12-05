package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"bubunyo/aoc-25/common"
)

const prefix = "day_5"

func main() {
	prefix := "Day 1:"
	fmt.Println(prefix, "Demo 1-", run_("input_demo"))
	fmt.Println(prefix, "Run 1 ->", run_("input"))
	fmt.Println(prefix, "Demo 2 ->", run2("input_demo"))
	fmt.Println(prefix, "Run 2 ->", run2("input"))
}

func run2(fp string) any {
	acc := 0
	d := map[int]map[int]struct{}{}
	pairs := true
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		str := strings.TrimSpace(str)
		if len(str) == 0 {
			pairs = false
			continue
		}
		if pairs {
			s := strings.Split(str, "|")
			l, err := strconv.Atoi(s[0])
			if err != nil {
				log.Fatalln("li", err)
			}
			r, err := strconv.Atoi(s[1])
			if err != nil {
				log.Fatalln("li", err)
			}
			set, ok := d[l]
			if ok {
				set[r] = struct{}{}
			} else {
				d[l] = map[int]struct{}{r: {}}
			}
		} else {
			rowStr := strings.Split(str, ",")
			row := make([]int, len(rowStr))
			for i, v := range rowStr {
				l, err := strconv.Atoi(v)
				if err != nil {
					log.Fatalln(err)
				}
				row[i] = l
			}

			valid := func() (bool, int, int) {
				for i := 0; i < len(row)-1; i++ {
					j := i + 1
					for ; j < len(row); j++ {
						dd := d[row[j]]
						if _, ok := dd[row[i]]; ok {
							return false, i, j
						}
					}
				}
				return true, -1, -1

			}

			v, i, j := valid()
			if v {
				continue
			}

			for !v {
				row[i], row[j] = row[j], row[i]
				v, i, j = valid()
			}
			// fmt.Println("valid>>", row, i, j, row[(len(row)-1)/2])
			acc += row[(len(row)-1)/2]
		}
	}
	return acc
}

func run_(fp string) any {
	acc := 0
	d := map[int]map[int]struct{}{}
	pairs := true
	for str := range common.IterateFileContent(prefix + "/" + fp) {
		str := strings.TrimSpace(str)
		if len(str) == 0 {
			pairs = false
			continue
		}
		if pairs {
			s := strings.Split(str, "|")
			l, err := strconv.Atoi(s[0])
			if err != nil {
				log.Fatalln("li", err)
			}
			r, err := strconv.Atoi(s[1])
			if err != nil {
				log.Fatalln("li", err)
			}
			set, ok := d[l]
			if ok {
				set[r] = struct{}{}
			} else {
				d[l] = map[int]struct{}{r: {}}
			}
		} else {
			rowStr := strings.Split(str, ",")
			row := make([]int, len(rowStr))
			for i, v := range rowStr {
				l, err := strconv.Atoi(v)
				if err != nil {
					log.Fatalln(err)
				}
				row[i] = l
			}

			valid := true
			for i := 0; i < len(row)-1; i++ {
				j := i + 1
				for ; j < len(row); j++ {
					dd := d[row[j]]
					if _, ok := dd[row[i]]; ok {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}

			if valid {
				acc += row[(len(row)-1)/2]
			}
		}
	}
	return acc
}

