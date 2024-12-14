package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
	"github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

var lenX, lenY int

type robot struct {
	position, velocity c2d.C
}

func (r *robot) move() {
	r.position.Add(r.velocity)
	if r.position.X >= lenX {
		r.position.X -= lenX
	}
	if r.position.X < 0 {
		r.position.X += lenX
	}
	if r.position.Y >= lenY {
		r.position.Y -= lenY
	}
	if r.position.Y < 0 {
		r.position.Y += lenY
	}
}

func RunSmall(input string) int {
	lenX = 11
	lenY = 7
	return Run(input)
}

func RunBig(input string) int {
	lenX = 101
	lenY = 103
	return Run(input)
}

func Run(input string) int {
	rows := common.ParseStringList(input, "\n")

	var robots []robot
	for _, row := range rows {
		r, _ := strings.CutPrefix(row, "p=")
		p := strings.Split(strings.Split(r, " v=")[0], ",")
		pos := c2d.C{X: common.MustAtoi(p[0]), Y: common.MustAtoi(p[1])}
		c := strings.Split(strings.Split(r, " v=")[1], ",")
		vel := c2d.C{X: common.MustAtoi(c[0]), Y: common.MustAtoi(c[1])}
		robots = append(robots, robot{pos, vel})
	}
	for i := 0; i < 100000000; i++ {
		for ri := range robots {
			robots[ri].move()
		}
		inTree := 0
		downMiddle := c2d.C{X: (lenX-1)/2, Y: lenY-1}
		for _, r := range robots {
			if downMiddle.ManhattanDistance(r.position) < lenY+2 {
				inTree++
			}
		}
		if inTree > 460 {
			print(i+1, robots)
			return i+1
		}
	}
	return 0
}

func print(iteration int, robots []robot) {
	fmt.Println("Iteration", iteration)
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if slices.ContainsFunc(robots, func(r robot) bool {
				return r.position.X == x && r.position.Y == y
			}) {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	submit := false
	verbose := true
	common.Run(RunBig, 2024, 14, 2, submit, verbose)
}
