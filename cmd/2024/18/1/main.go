package main

import (
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

func Run70(input string) int {
	return Run(input, C{70, 70}, 1024)
}

var dirs []C = []C{Up, Down, Right, Left}

func Run(input string, end C, fallen int) int {
	rows := common.ParseStringList(input, "\n")
	var falling []C
	for i, r := range rows {
		if i >= fallen {
			break
		}
		falling = append(falling, C{common.MustAtoi(strings.Split(r, ",")[0]), common.MustAtoi(strings.Split(r, ",")[1])})
	}
	visited := map[C]int{C{0, 0}: 0}
	current := []C{C{0, 0}}
	t := 0
	for len(current) > 0 {
		t++
		var next []C
		for _, c := range current {
			for _, d := range dirs {
				nxt := c.Plus(d)
				_, isVisited := visited[nxt]
				isCorrupted := slices.Contains(falling, nxt)
				if nxt.X >= 0 && nxt.X <= end.X && nxt.Y >= 0 && nxt.Y <= end.Y && !isVisited && !isCorrupted {
					next = append(next, nxt)
					visited[nxt] = t
					if nxt == end {
						return t
					}
				}
			}
		}
		current = next
	}
	return 0
}

func main() {
	submit := false
	verbose := true
	common.Run(Run70, 2024, 18, 1, submit, verbose)
}
