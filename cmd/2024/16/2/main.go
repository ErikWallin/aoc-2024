package main

import (
	"fmt"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

type Score struct {
	value     int
	position  C
	direction C
	paths     map[C]bool
}

func (s Score) key() string {
	return key(s.position, s.direction)
}

func key(position, direction C) string {
	return fmt.Sprintf("%d,%d+%d,%d", position.X, position.Y, direction.X, direction.Y)
}

func Run(input string) int {
	runes := common.ParseRuneListList(input)
	var start, end C
	for y, row := range runes {
		for x, r := range row {
			if r == 'S' {
				start = C{x, y}
			} else if r == 'E' {
				end = C{x, y}
			}
		}
	}
	records := map[string]Score{}
	startScore := Score{0, start, Right, map[C]bool{start: true}}
	currentRound := []Score{startScore}
	records[startScore.key()] = startScore
	for len(currentRound) > 0 {
		var nextRound []Score
		for _, current := range currentRound {
			nextPosition := current.position.Plus(current.direction)
			if runes[nextPosition.Y][nextPosition.X] != '#' {
				nextPaths := map[C]bool{}
				for previousPathItem := range current.paths {
					nextPaths[previousPathItem] = true
				}
				nextPaths[nextPosition] = true
				next := Score{current.value + 1, nextPosition, current.direction, nextPaths}
				previousScore, ok := records[next.key()]
				if !ok || next.value < previousScore.value {
					records[next.key()] = next
					nextRound = append(nextRound, next)
				} else if !ok || next.value == previousScore.value {
					for pp := range previousScore.paths {
						next.paths[pp] = true
					}
					next.paths[next.position] = true
					nextRound = append(nextRound, next)
				}
			}
			if current.direction == Up || current.direction == Down {
				nextPaths := map[C]bool{}
				for previousPathItem := range current.paths {
					nextPaths[previousPathItem] = true
				}
				next := Score{current.value + 1000, current.position, Left, nextPaths}
				previousScore, ok := records[next.key()]
				if !ok || next.value < previousScore.value {
					records[next.key()] = next
					nextRound = append(nextRound, next)
				} else if !ok || next.value == previousScore.value {
					for previousPathItem := range previousScore.paths {
						next.paths[previousPathItem] = true
					}
					records[next.key()] = next
					nextRound = append(nextRound, next)
				}
				nextPaths = map[C]bool{}
				for previousPathItem := range current.paths {
					nextPaths[previousPathItem] = true
				}
				next = Score{current.value + 1000, current.position, Right, nextPaths}
				previousScore, ok = records[next.key()]
				if !ok || next.value < previousScore.value {
					records[next.key()] = next
					nextRound = append(nextRound, next)
				} else if !ok || next.value == previousScore.value {
					for pp := range previousScore.paths {
						next.paths[pp] = true
					}
					records[next.key()] = next
					nextRound = append(nextRound, next)
				}
			}
			if current.direction == Left || current.direction == Right {
				path := map[C]bool{}
				for previousPathItem := range current.paths {
					path[previousPathItem] = true
				}
				sn := Score{current.value + 1000, current.position, Up, path}
				pc, ok := records[sn.key()]
				if !ok || sn.value < pc.value {
					records[sn.key()] = sn
					nextRound = append(nextRound, sn)
				} else if !ok || sn.value == pc.value {
					for previousPathItem := range current.paths {
						sn.paths[previousPathItem] = true
					}
					records[sn.key()] = sn
					nextRound = append(nextRound, sn)
				}
				path = map[C]bool{}
				for pp := range current.paths {
					path[pp] = true
				}
				sn = Score{current.value + 1000, current.position, Down, path}
				pc, ok = records[sn.key()]
				if !ok || sn.value < pc.value {
					records[sn.key()] = sn
					nextRound = append(nextRound, sn)
				} else if !ok || sn.value == pc.value {
					for pp := range pc.paths {
						sn.paths[pp] = true
					}
					records[sn.key()] = sn
					nextRound = append(nextRound, sn)
				}
			}
		}
		currentRound = nextRound
	}
	possibleEndKeys := []string{key(end, Up), key(end, Down), key(end, Right), key(end, Left)}
	minKey := ""
	min := -1
	for _, k := range possibleEndKeys {
		if val, ok := records[k]; ok {
			if min == -1 || val.value < min {
				minKey = k
			}
		}
	}
	return len(records[minKey].paths)
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 16, 2, submit, verbose)
}
