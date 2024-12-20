package main

import (
	"fmt"
	"slices"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

var ds = map[rune]C{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

func Run(input string) int {
	sections := common.ParseStringList(input, "\n\n")
	pool := common.ParseRuneListList(sections[0])
	movements := sections[1]
	var pos C
	for y, row := range pool {
		for x, r := range row {
			if r == '@' {
				pos = C{x, y}
			}
		}
	}
	valid := []rune{'<', '>', '^', 'v'}
	for _, m := range movements {
		if !slices.Contains(valid, m) {
			continue
		}
		dir := ds[m]
		to := pos
		next := pos
		for {
			next.Add(dir)
			if pool[next.Y][next.X] == '#' {
				break
			}
			if pool[next.Y][next.X] == '.' {
				to = next
				break
			}
		}
		if to.X == pos.X && to.Y == pos.Y {
			continue
		}
		for {
			if pool[to.Y][to.X] == '.' {
				prev := to.Minus(dir)
				pool[to.Y][to.X] = pool[prev.Y][prev.X]
				if pool[prev.Y][prev.X] == '@' {
					pos = to
				}
				pool[prev.Y][prev.X] = '.'
			}
			if to.X == pos.X && to.Y == pos.Y {
				break
			}
			to.Subtract(dir)
		}
	}

	print(pool)
	sum := 0
	for y, row := range pool {
		for x, r := range row {
			if r == 'O' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func print(pool [][]rune) {
	for _, row := range pool {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 15, 1, submit, verbose)
}
