package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	lines := common.ParseStringList(input, "\n")
	count := 0
	for _, l := range lines {
		sections := common.ParseStringList(l, ": ")
		expected := common.MustAtoi(sections[0])
		numbers := common.ParseIntList(sections[1], " ")
		res := operate(numbers[0], numbers[1:])
		for _, r := range res {
			if r == expected {
				count += expected
				break
			}
		}
	}
	return count
}

func operate(value int, rest []int) []int {
	if len(rest) == 0 {
		return []int{value}
	}
	var res []int
	add := operate(value + rest[0], rest[1:])
	res = append(res, add...)
	multiply := operate(value * rest[0], rest[1:])
	res = append(res, multiply...)
	return res
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 7, 1, submit, verbose)
}
