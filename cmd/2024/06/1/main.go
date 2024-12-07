package main

import (
	"fmt"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

type C struct {
	x int
	y int
}

var up = C{0, -1}
var down = C{0, 1}
var right = C{1, 0}
var left = C{-1, 0}

func (c1 C) add(c2 C) C {
	return C{c1.x + c2.x, c1.y + c2.y}
}

func (c C) turnRight() C {
	if c == up {
		return right
	} else if c == right {
		return down
	} else if c == down {
		return left
	} else if c == left {
		return up
	}
	panic(c)
}

func Run(input string) int {
	up := C{0, -1}
	down := C{0, 1}
	right := C{1, 0}
	left := C{-1, 0}
	visitedDir := map[string]bool{}
	visited := map[C]bool{}
	var pos C
	var dir C
	field := common.ParseRuneListList(input)
	for y, r := range field {
		for x, c := range r {
			if c == '^' {
				pos = C{x, y}
				dir = up
			} else if c == 'v' {
				pos = C{x, y}
				dir = down
			} else if c == '<' {
				pos = C{x, y}
				dir = left
			} else if c == '>' {
				pos = C{x, y}
				dir = right
			}
		}
	}
	for {
		if _, ok := visitedDir[visitedKey(pos, dir)]; ok {
			break
		}
		visitedDir[visitedKey(pos, dir)] = true
		visited[pos] = true
		next := pos.add(dir)
		if next.x < 0 || next.y < 0 || next.x >= len(field[0]) || next.y >= len(field) {
			break
		}
		if field[next.y][next.x] != '#' {
			pos = next
			continue
		}
		dir = dir.turnRight()
		next = pos.add(dir)
		if next.x < 0 || next.y < 0 || next.x >= len(field[0]) || next.y >= len(field) {
			break
		}
		if field[next.y][next.x] != '#' {
			pos = next
			continue
		}
		dir = dir.turnRight()
		next = pos.add(dir)
		if next.x < 0 || next.y < 0 || next.x >= len(field[0]) || next.y >= len(field) {
			break
		}
		if field[next.y][next.x] != '#' {
			pos = next
			continue
		}
		dir = dir.turnRight()
		next = pos.add(dir)
		if next.x < 0 || next.y < 0 || next.x >= len(field[0]) || next.y >= len(field) {
			break
		}
		if field[next.y][next.x] != '#' {
			pos = next
			continue
		}
		break
	}

	return len(visited)
}

func visitedKey(pos, dir C) string {
	return fmt.Sprintf("%v|%v", pos, dir)
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 6, 1, submit, verbose)
}
