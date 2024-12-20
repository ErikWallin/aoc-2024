package main

import (
	"slices"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

var dir []C = []C{Up, Right, Down, Left}

func Run100(input string) int {
	return Run(input, 100)
}

func Run(input string, save int) int {
	f := common.ParseRuneListList(input)

	var start, end C
	for y, row := range f {
		for x, r := range row {
			if r == 'S' {
				start = C{x, y}
			}
			if r == 'E' {
				end = C{x, y}
			}
		}
	}

	trackWithoutCheat := []C{start}
	current := start
	t := 0
	for current != end {
		t++
		for _, d := range dir {
			next := current.Plus(d)
			if slices.Contains(trackWithoutCheat, next) || next.X < 0 || next.Y < 0 || next.X >= len(f[0]) || next.Y >= len(f) || f[next.Y][next.X] == '#' {
				continue
			}
			trackWithoutCheat = append(trackWithoutCheat, next)
			current = next
			break
		}
	}

	count := 0
	for t1, c1 := range trackWithoutCheat {
		for t2, c2 := range trackWithoutCheat[t1+1:] {
			if c1.ManhattanDistance(c2) <= 20 && t2+1 - c1.ManhattanDistance(c2) >= save {
				count++
			}
		}
	}
	return count
}

func main() {
	submit := false
	verbose := true
	common.Run(Run100, 2024, 20, 2, submit, verbose)
}
