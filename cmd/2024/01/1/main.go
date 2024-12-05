package main

import (
	"slices"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	inp := common.ParseIntListList(input, "   ")
	var xs, ys []int
	for _, line := range inp {
		xs = append(xs, line[0])
		ys = append(ys, line[1])
	}
	slices.Sort(xs)
	slices.Sort(ys)
	var sumDistance int
	for i := range len(inp) {
		d := xs[i] - ys[i]
		if d < 0 {
			d = -d
		}
		sumDistance += d
	}
	return sumDistance
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 1, 1, submit, verbose)
}
