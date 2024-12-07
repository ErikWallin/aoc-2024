package main

import (
	"strconv"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	rest := input
	sum := 0
	isDo := true
	for {
		if len(rest) < 8 {
			break
		}
		if strings.HasPrefix(rest, "do()") {
			isDo = true
			rest = rest[4:]
		}
		if strings.HasPrefix(rest, "don't()") {
			isDo = false
			rest = rest[7:]
		}
		if !strings.HasPrefix(rest, "mul(") {
			rest = rest[1:]
			continue
		}
		rest = rest[4:]
		commaI := strings.Index(rest, ",")
		n1, err1 := strconv.Atoi(rest[:commaI])
		if err1 != nil {
			continue
		}
		rest = rest[commaI+1:]
		parenthesisI := strings.Index(rest, ")")
		n2, err2 := strconv.Atoi(rest[:parenthesisI])
		if err2 != nil {
			continue
		}
		if isDo {
			sum += n1 * n2
		}
	}
	return sum
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 3, 2, submit, verbose)
}
