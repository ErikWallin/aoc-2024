package main

import (
	"math"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
	"github.com/ErikWallin/aoc-2024/internal/common/c2d"
	"gonum.org/v1/gonum/mat"
)

func parseCoordinate(row, prefix, sep string) c2d.C {
	s, _ := strings.CutPrefix(row, prefix)
	xy := strings.Split(s, sep)
	return c2d.C{X: common.MustAtoi(xy[0]), Y: common.MustAtoi(xy[1])}
}

func RunWithPrizeOffset(input string) int {
	return Run(input, 10000000000000)
}

func Run(input string, prizeOffset int) int {
	sections := common.ParseStringList(input, "\n\n")
	sum := 0
	for _, s := range sections {
		rows := common.ParseStringList(s, "\n")
		buttonA := parseCoordinate(rows[0], "Button A: X+", ", Y+")
		buttonB := parseCoordinate(rows[1], "Button B: X+", ", Y+")
		prize := parseCoordinate(rows[2], "Prize: X=", ", Y=")

		matrix := mat.NewDense(2, 2, []float64{float64(buttonA.X), float64(buttonB.X), float64(buttonA.Y), float64(buttonB.Y)})
		solutionVector := mat.NewDense(2, 1, []float64{float64(prize.X), float64(prize.Y)})
		sol := mat.NewDense(2, 1, []float64{0, 0})
		sol.Solve(matrix, solutionVector)
		pressedA := int(math.Round(sol.At(0, 0)))
		pressedB := int(math.Round(sol.At(1, 0)))
		if pressedA*buttonA.X+pressedB*buttonB.X == prize.X && pressedA*buttonA.Y+pressedB*buttonB.Y == prize.Y {
			sum += 3*pressedA + pressedB
		}
	}
	return sum
}

func main() {
	submit := false
	verbose := true
	common.Run(RunWithPrizeOffset, 2024, 13, 2, submit, verbose)
}
