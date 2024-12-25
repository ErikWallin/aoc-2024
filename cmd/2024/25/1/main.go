package main

import (
	"slices"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	locks := [][]int{}
	keys := [][]int{}
	parts := common.ParseStringList(input, "\n\n")
	for _, part := range parts {
		lockOrKey, isLock := getLockOrKey(part)
		if isLock {
			locks = append(locks, lockOrKey)
		} else {
			keys = append(keys, lockOrKey)
		}
	}

	sum := 0
	for _, k := range keys {
		for _, l := range locks {
			overlapping := false
			for x := range len(l) {
				if k[x]+l[x] > 5 {
					overlapping = true
					break
				}
			}
			if !overlapping {
				sum++
			}
		}
	}
	return sum
}

func getLockOrKey(part string) ([]int, bool) {
	rows := common.ParseStringList(part, "\n")
	isLock := true
	if rows[0][0] != '#' {
		isLock = false
		slices.Reverse(rows)
	}
	values := []int{}
	for x := range len(rows[0]) {
		for y := range len(rows) - 1 {
			if rows[y+1][x] == '.' {
				values = append(values, y)
				break
			}
		}
	}
	return values, isLock
}

func main() {
	submit := true
	verbose := true
	common.Run(Run, 2024, 25, 1, submit, verbose)
}
