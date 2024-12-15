package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

type C struct {
	x, y int
}

func (c C) neighbours(lenX, lenY int) []C {
	var ns []C
	if c.x > 0 {
		ns = append(ns, C{c.x - 1, c.y})
	}
	if c.x < lenX-1 {
		ns = append(ns, C{c.x + 1, c.y})
	}
	if c.y > 0 {
		ns = append(ns, C{c.x, c.y - 1})
	}
	if c.y < lenY-1 {
		ns = append(ns, C{c.x, c.y + 1})
	}
	return ns
}

func Run(input string) int {
	digits := common.ParseDigitListList(input)
	lenX := len(digits[0])
	lenY := len(digits)
	var trailheads []C
	for y := range len(digits) {
		for x := range len(digits[0]) {
			if digits[y][x] == 0 {
				trailheads = append(trailheads, C{x, y})
			}
		}
	}
	sum := 0

	for _, t := range trailheads {
		sum += traverse(t, digits, lenX, lenY)
	}

	return sum
}

func traverse(t C, digits [][]int, lenX, lenY int) int {
	currentPos := []C{t}
	var nextPos []C
	for next := 1; next <= 9; next++ {
		nextPos = []C{}
		for _, cp := range currentPos {
			for _, n := range cp.neighbours(lenX, lenY) {
				if digits[n.y][n.x] == next {
					nextPos = append(nextPos, n)
				}
			}
		}
		currentPos = nextPos
	}
	return len(nextPos)
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 10, 1, submit, verbose)
}
