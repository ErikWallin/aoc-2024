package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

type C struct {
	x int
	y int
}

func Run(input string) int {
	puzzle := common.ParseStringList(input, "\n")
	xLen := len(puzzle[0])
	yLen := len(puzzle)
	var starts []C
	for y, row := range puzzle {
		for x, c := range row {
			if c == 'A' {
				starts = append(starts, C{x: x, y: y})
			}
		}
	}
	count := 0
	for _, start := range starts {
		if start.x + 1 < xLen && start.x - 1 >= 0 && start.y + 1 < yLen && start.y - 1 >= 0 {
			// Down Right
			c1 := puzzle[start.y+1][start.x+1]
			c2 := puzzle[start.y-1][start.x-1]
			if !((c1 == 'M' && c2 == 'S') || (c1 == 'S' && c2 == 'M')) {
				continue
			}
			// Up Right
			c1 = puzzle[start.y-1][start.x+1]
			c2 = puzzle[start.y+1][start.x-1]
			if !((c1 == 'M' && c2 == 'S') || (c1 == 'S' && c2 == 'M')) {
				continue
			}
			count++
		}
	}

	return count
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 4, 2, submit, verbose)
}
