package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

var rules [][]int

func Run(input string) int {
	sections := strings.Split(input, "\n\n")
	for _, r := range common.ParseStringList(sections[0], "\n") {
		rules = append(rules, common.ParseIntList(r, "|"))
	}
	var updates [][]int
	for _, r := range common.ParseStringList(sections[1], "\n") {
		updates = append(updates, common.ParseIntList(r, ","))
	}
	fmt.Println(rules)
	fmt.Println(updates)

	var incorrectUpdates [][]int
	for _, update := range updates {
		isOk := true
		for _, rule := range rules {
			if !isUpdateOk(update, rule) {
				isOk = false
				break
			}
		}
		if !isOk {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}
	sum := 0
	for _, update := range incorrectUpdates {
		slices.SortFunc(update, compare)
		sum += update[len(update) / 2]
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

func compare(a, b int) int {
	for _, r := range rules {
		if r[0] == a && r[1] == b {
			return -1
		} else if r[1] == a && r[0] == b {
			return 1
		}
	}
	return 0
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 5, 1, submit, verbose)
}
