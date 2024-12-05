package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	inp := common.ParseIntListList(input, "   ")
	var xs, ys []int
	for _, line := range inp {
		xs = append(xs, line[0])
		ys = append(ys, line[1])
	}
	var similarityScore int
	for _, x := range xs {
		var times int
		for _, y := range ys {
			if x == y {
				times++
			}
		}
		similarityScore += x * times
	}
	return similarityScore
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 1, 2, submit, verbose)
}
