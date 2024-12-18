package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

func Run70(input string) int {
	return Run(input, C{70, 70})
}

var dirs []C = []C{Up, Down, Right, Left}

var result string

func Run(input string, end C) int {
	rows := common.ParseStringList(input, "\n")
	var falling []C
	for _, r := range rows {
		falling = append(falling, C{common.MustAtoi(strings.Split(r, ",")[0]), common.MustAtoi(strings.Split(r, ",")[1])})
	}
	var corrupted []C
	f := 382
	for {
		f++
		visited := map[C]int{C{0, 0}: 0}
		current := []C{C{0, 0}}
		if f >= len(falling) {
			break
		}
		corrupted = falling[:f]
		t := 0
		success := false
		for len(current) > 0 {
			t++
			var next []C
			for _, c := range current {
				for _, d := range dirs {
					nxt := c.Plus(d)
					_, isVisited := visited[nxt]
					isCorrupted := slices.Contains(corrupted, nxt)
					if nxt.X >= 0 && nxt.X <= end.X && nxt.Y >= 0 && nxt.Y <= end.Y && !isVisited && !isCorrupted {
						next = append(next, nxt)
						visited[nxt] = t
						if nxt == end {
							success = true
							break
						}
					}
				}
				if success {
					break
				}
			}
			current = next
		}
		if !success {
			lastCorrupted := corrupted[len(corrupted) - 1]
			result = fmt.Sprintf("%d,%d", lastCorrupted.X, lastCorrupted.Y)
			break
		}
	}
	
	fmt.Println("Result", result)
	return 0
}

func main() {
	submit := false
	verbose := true
	common.Run(Run70, 2024, 18, 2, submit, verbose)
}
