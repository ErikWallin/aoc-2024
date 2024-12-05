package main

import (
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)



func Run(input string) int {
	sections := strings.Split(input, "\n\n")
	var rules [][]int
	for _, r := range common.ParseStringList(sections[0], "\n") {
		rules = append(rules, common.ParseIntList(r, "|"))
	}
	var updates [][]int
	for _, r := range common.ParseStringList(sections[1], "\n") {
		updates = append(updates, common.ParseIntList(r, ","))
	}

	sum := 0
	for _, update := range updates {
		isOk := true
		for _, rule := range rules {
			if !isUpdateOk(update, rule) {
				isOk = false
				break
			}
		}
		if isOk {
			sum += update[len(update) / 2]
		}
	}
	return sum
}

func isUpdateOk(update []int, rule []int) bool {
	i1 := slices.Index(update, rule[0])
	i2 := slices.Index(update, rule[1])
	if i1 == -1 || i2 == -1 || i1 < i2 {
		return true
	}
	return false
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 5, 1, submit, verbose)
}
