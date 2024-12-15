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
			if c == 'X' {
				starts = append(starts, C{x: x, y: y})
			}
		}
	}
	count := 0
	for _, start := range starts {
		// Right
		if start.x+3 < xLen {
			if puzzle[start.y][start.x+1] == 'M' && puzzle[start.y][start.x+2] == 'A' && puzzle[start.y][start.x+3] == 'S' {
				count++
			}
		}
		// Left
		if start.x-3 >= 0 {
			if puzzle[start.y][start.x-1] == 'M' && puzzle[start.y][start.x-2] == 'A' && puzzle[start.y][start.x-3] == 'S' {
				count++
			}
		}
		// Down
		if start.y+3 < yLen {
			if puzzle[start.y+1][start.x] == 'M' && puzzle[start.y+2][start.x] == 'A' && puzzle[start.y+3][start.x] == 'S' {
				count++
			}
		}
		// Up
		if start.y-3 >= 0 {
			if puzzle[start.y-1][start.x] == 'M' && puzzle[start.y-2][start.x] == 'A' && puzzle[start.y-3][start.x] == 'S' {
				count++
			}
		}
		// Down Right
		if start.x+3 < xLen && start.y+3 < yLen {
			if puzzle[start.y+1][start.x+1] == 'M' && puzzle[start.y+2][start.x+2] == 'A' && puzzle[start.y+3][start.x+3] == 'S' {
				count++
			}
		}
		// Up Right
		if start.x+3 < xLen && start.y-3 >= 0 {
			if puzzle[start.y-1][start.x+1] == 'M' && puzzle[start.y-2][start.x+2] == 'A' && puzzle[start.y-3][start.x+3] == 'S' {
				count++
			}
		}
		// Up Left
		if start.x-3 >= 0 && start.y-3 >= 0 {
			if puzzle[start.y-1][start.x-1] == 'M' && puzzle[start.y-2][start.x-2] == 'A' && puzzle[start.y-3][start.x-3] == 'S' {
				count++
			}
		}
		// Down Left
		if start.x-3 >= 0 && start.y+3 < yLen {
			if puzzle[start.y+1][start.x-1] == 'M' && puzzle[start.y+2][start.x-2] == 'A' && puzzle[start.y+3][start.x-3] == 'S' {
				count++
			}
		}
	}

	return count
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 4, 1, submit, verbose)
}
