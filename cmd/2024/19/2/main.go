package main

import (
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

var towels []string

func Run(input string) int {
	sections := common.ParseStringList(input, "\n\n")
	towels = common.ParseStringList(sections[0], ", ")
	patterns := common.ParseStringList(sections[1], "\n")
	possible := 0
	for _, pattern := range patterns {
		possible += countPossible(pattern)
	}
	return int(possible)
}

var cache map[string]int = map[string]int{}

func countPossible(pattern string) int {
	if cachedValue, ok := cache[pattern]; ok {
		return cachedValue
	}
	if pattern == "" {
		return 1
	}
	sum := 0
	for _, towel := range towels {
		if after, found := strings.CutPrefix(pattern, towel); found {
			sum += countPossible(after)
		}
	}
	cache[pattern] = sum
	return sum
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 19, 2, submit, verbose)
}
