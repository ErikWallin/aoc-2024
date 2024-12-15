package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	inp := common.ParseIntListList(input, " ")
	var count int
	for _, r := range inp {
		if isDecreasing(r) {
			count++
		}
	}
	for _, r := range inp {
		if isIncreasing(r) {
			count++
		}
	}
	return count
}

func isIncreasing(r []int) bool {
	current := r[0] - 1
	for _, i := range r {
		if i <= current || i > current+3 {
			return false
		}
		current = i
	}
	return true
}

func isDecreasing(r []int) bool {
	current := r[0] + 1
	for _, i := range r {
		if i >= current || i < current-3 {
			return false
		}
		current = i
	}
	return true
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 2, 1, submit, verbose)
}
