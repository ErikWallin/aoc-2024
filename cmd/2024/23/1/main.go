package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

var connections map[string][]string = map[string][]string{}
var ts []string = []string{}

func Run(input string) int {
	rows := common.ParseStringList(input, "\n")
	for _, row := range rows {
		parts := strings.Split(row, "-")
		//fmt.Println(parts)
		first := parts[0]
		second := parts[1]
		if _, ok := strings.CutPrefix(first, "t"); ok {
			if !slices.Contains(ts, first) {
				ts = append(ts, first)
			}
		}
		if _, ok := strings.CutPrefix(second, "t"); ok {
			if !slices.Contains(ts, second) {
				ts = append(ts, second)
			}
		}
		if _, found := connections[first]; !found {
			connections[first] = []string{}
		}
		if _, found := connections[second]; !found {
			connections[second] = []string{}
		}
		connections[first] = append(connections[first], second)
		connections[second] = append(connections[second], first)
	}
	sum := 0
	for it, t := range ts {
		for ic1, c1 := range connections[t] {
			if _, ok := strings.CutPrefix(c1, "t"); ok {
				if slices.Index(ts, c1) < it {
					continue
				}
			}
			for ic2, c2 := range connections[t] {
				if _, ok := strings.CutPrefix(c2, "t"); ok {
					if slices.Index(ts, c2) < it {
						continue
					}
				}
				if ic2 > ic1 && slices.Contains(connections[c1], c2) {
					sum++
					fmt.Println(t, c1, c2)
				}
			}
		}
	}
	return sum
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 23, 1, submit, verbose)
}
