package main

import (
	"math"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	ints := common.ParseIntList(input, " ")
	sum := 0
	for _, i := range ints {
		sum += len(BlinkNTimes(i, 25))
	}
	return sum
}

func BlinkNTimes(i, times int) []int {
	if times == 0 {
		return []int{i}
	}
	b := blink(i)
	if len(b) == 1 {
		return BlinkNTimes(b[0], times - 1)
	} else {
		res := BlinkNTimes(b[0], times - 1)
		res = append(res, BlinkNTimes(b[1], times - 1)...)
		return res
	}
}

func blink(i int) []int {
	if i == 0 {
		return []int{1}
	}
	l := lenDigit(i)
	if l % 2 == 0 {
		a := int(math.Pow10(l/2))
		return []int{i / a, i % a}
	}
	return []int{i * 2024}
}

func lenDigit(i int) int {
    if i == 0 {
        return 1
    }
    count := 0
    for i != 0 {
        i /= 10
        count++
    }
    return count
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 11, 1, submit, verbose)
}
