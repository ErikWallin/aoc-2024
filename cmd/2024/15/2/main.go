package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

var ds = map[rune]C{
	'^': C{0, -1},
	'v': C{0, 1},
	'<': C{-1, 0},
	'>': C{1, 0},
}

var pool [][]rune

func Run(input string) int {
	sections := common.ParseStringList(input, "\n\n")
	for _, row := range strings.Split(sections[0], "\n") {
		newRow := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(row, "#", "##"), "O", "[]"), ".", ".."), "@", "@.")
		pool = append(pool, []rune(newRow))
	}
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
		if canPush(pos, dir) {
			push(pos, dir)
			pool[pos.Y][pos.X] = '.'
			pos.Add(dir)
		}
	}
	print()
	sum := 0
	for y, row := range pool {
		for x, r := range row {
			if r == '[' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func canPush(pos, dir C) bool {
	next := pos.Plus(dir)
	if pool[next.Y][next.X] == '.' {
		return true
	}
	if pool[next.Y][next.X] == '#' {
		return false
	}
	if dir.Y == 0 {
		return canPush(pos.Plus(dir), dir)
	} else {
		if pool[next.Y][next.X] == '[' {
			return canPush(pos.Plus(dir), dir) && canPush(pos.Plus(dir).Plus(C{1, 0}), dir)
		} else {
			return canPush(pos.Plus(dir), dir) && canPush(pos.Plus(dir).Plus(C{-1, 0}), dir)
		}
	}
}

func push(pos, dir C) {
	if pool[pos.Y][pos.X] == '.' {
		return
	}
	next := pos.Plus(dir)
	if dir.X == 0 {
		if pool[next.Y][next.X] == '[' {
			neighbour := next.Plus(C{1, 0})
			push(next, dir)
			push(neighbour, dir)
			pool[next.Y][next.X] = pool[pos.Y][pos.X]
			pool[neighbour.Y][neighbour.X] = '.'
		} else if pool[next.Y][next.X] == ']' {
			neighbour := next.Plus(C{-1, 0})
			push(next, dir)
			push(neighbour, dir)
			pool[next.Y][next.X] = pool[pos.Y][pos.X]
			pool[neighbour.Y][neighbour.X] = '.'
		} else {
			push(next, dir)
			pool[next.Y][next.X] = pool[pos.Y][pos.X]
		}
	} else if dir.Y == 0 {
		push(next, dir)
			pool[next.Y][next.X] = pool[pos.Y][pos.X]
	}
}

func print() {
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
	common.Run(Run, 2024, 15, 2, submit, verbose)
}
