package main

import (
	"slices"
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
		if isPossible(pattern) {
			possible++
		}
	}
	return possible
}

func isPossible(pattern string) bool {
	if slices.Contains(towels, pattern) {
		return true
	}
	for _, towel := range towels {
		if after, found := strings.CutPrefix(pattern, towel); found {
			if isPossible(after) {
				return true
			}
		}
	}
	return false
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 19, 1, submit, verbose)
}
