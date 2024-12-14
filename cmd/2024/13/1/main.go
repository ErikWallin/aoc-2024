package main

import (
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
	"github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

func parseCoordinate(row, prefix, sep string) c2d.C {
	s, _ := strings.CutPrefix(row, prefix)
	xy := strings.Split(s, sep)
	return c2d.C{X: common.MustAtoi(xy[0]), Y: common.MustAtoi(xy[1])}
}

func Run(input string) int {
	sections := common.ParseStringList(input, "\n\n")
	sum := 0
	for _, s := range sections {
		rows := common.ParseStringList(s, "\n")
		buttonA := parseCoordinate(rows[0], "Button A: X+", ", Y+")
		buttonB := parseCoordinate(rows[1], "Button B: X+", ", Y+")
		prize := parseCoordinate(rows[2], "Prize: X=", ", Y=")

		var pressedA, pressedB int
		success := false
		for pa := range 101 {
			pressedA = pa - 1
			for pb := range 101 {
				pressedB = pb - 1
				if pressedA*buttonA.X+pressedB*buttonB.X == prize.X && pressedA*buttonA.Y+pressedB*buttonB.Y == prize.Y {
					success = true
					break
				}
			}
			if success {
				break
			}
		}
		if success {
			sum += 3*pressedA + pressedB
		}
	}
	return sum
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 13, 1, submit, verbose)
}
