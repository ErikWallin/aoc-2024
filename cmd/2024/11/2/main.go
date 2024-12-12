package main

import (
	"math"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run75(input string) int {
	return Run(input, 75)
}

func Run(input string, times int) int {
	ints := common.ParseIntList(input, " ")
	sum := 0
	cache = map[int]int{}
	for _, i := range ints {
		sum += blinkNTimes(i, times)
	}
	return sum
}

var cache map[int]int

func cacheKey(i, times int) int {
	return i * 100 + times
}

func blinkNTimes(i, times int) int {
	if times == 0 {
		return 1
	}
	ck := cacheKey(i, times)
	if val, ok := cache[ck]; ok {
		return val
	}
	b := blink(i)
	if len(b) == 1 {
		res := blinkNTimes(b[0], times - 1)
		cache[ck] = res
		return res
	} else {
		res := blinkNTimes(b[0], times - 1) + blinkNTimes(b[1], times - 1)
		cache[ck] = res
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
	common.Run(Run75, 2024, 11, 1, submit, verbose)
}
