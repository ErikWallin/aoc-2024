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
			continue
		}
		isSafe := false
		for ri, _ := range r {
			var rMod []int
			for i, _ := range r {
				if i != ri {
					rMod = append(rMod, r[i])
				}
			}
			if isDecreasing(rMod) {
				isSafe = true
			}
		}
		if isSafe {
			count++
			continue
		}
		if isIncreasing(r) {
			count++
			continue
		}
		for ri, _ := range r {
			var rMod []int
			for i, _ := range r {
				if i != ri {
					rMod = append(rMod, r[i])
				}
			}
			if isIncreasing(rMod) {
				isSafe = true
			}
		}
		if isSafe {
			count++
			continue
		}
	}
	return count
}

func isIncreasing(r []int) bool {
	current := r[0]-1
	for _, i := range r {
		if i <= current || i > current+3  {
			return false
		}
		current = i
	}
	return true
}

func isDecreasing(r []int) bool {
	current := r[0]+1
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
