package main

import (
	"fmt"

	"github.com/ErikWallin/aoc-2024/internal/common"
	"github.com/samber/lo"
)

func Run(input string) int {
	calibration_document := common.ParseRuneListList(input)
	numbers := []int{}
	for _, line := range calibration_document {
		var first, last rune
		for _, r := range line {
			if r >= '0' && r <= '9' {
				if first == 0 {
					first = r
				}
				last = r
			}
		}
		number := fmt.Sprintf("%v%v", string(first), string(last))
		numbers = append(numbers, common.MustAtoi(number))
	}
	return lo.Sum(numbers)
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 1, 1, submit, verbose)
}
