package common

import (
	"strings"
)

// Parse input to slice of integers, i.e. []int
// Works for both csv (sep = ",") and row delimited (sep = "\n")
func ParseIntList(input string, sep string) []int {
	lines := strings.Split(input, sep)
	list := make([]int, len(lines))
	for i, line := range lines {
		list[i] = MustAtoi(line)
	}
	return list
}

func ParseDigitList(input string) []int {
	list := make([]int, len(input))
	for i, r := range input {
		list[i] = MustAtoi(string(r))
	}
	return list
}

// Parse input to slice of strings, i.e. []string
// Works for both csv (sep = ",") and row delimited (sep = "\n")
func ParseStringList(input string, sep string) []string {
	lines := strings.Split(input, sep)
	return lines
}

func ParseRuneListList(input string) [][]rune {
	var xys [][]rune
	for _, x := range strings.Split(input, "\n") {
		xys = append(xys, []rune(x))
	}
	return xys
}

func ParseDigitListList(input string) [][]int {
	rows := strings.Split(input, "\n")
	ys := make([][]int, len(rows))
	for y, row := range rows {
		ys[y] = make([]int, len(row))
		for x, r := range row {
			ys[y][x] = MustAtoi(string(r))
		}
	}
	return ys
}

func ParseRuneListListList(input string) [][][]rune {
	var xyzs [][][]rune
	var yzs [][]rune
	for _, x := range strings.Split(input, "\n") {
		if len(x) == 0 {
			xyzs = append(xyzs, yzs)
			yzs = nil
		} else {
			yzs = append(yzs, []rune(x))
		}
	}
	xyzs = append(xyzs, yzs)
	return xyzs
}

func ParseIntListList(input string, sep string) [][]int {
	var xys [][]int
	for _, line := range strings.Split(input, "\n") {
		var xs []int
		for _, x := range strings.Split(line, sep) {
			xs = append(xs, MustAtoi(x))
		}
		xys = append(xys, xs)
	}
	return xys
}
