package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 3
	got := Run(`#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}

func TestGetLock(t *testing.T) {
	expected := []int{0, 5, 3, 4, 3}
	got, isLock := getLockOrKey(`#####
.####
.####
.####
.#.#.
.#...
.....`)
	if len(got) != len(expected) {
		t.Errorf("GetLock = %d; expected %d", got, expected)
	}
	if !isLock {
		t.Errorf("GetLock isLock = %t; expected %t", isLock, true)
	}
	for x, v := range expected {
		if got[x] != v {
			t.Errorf("Run = %d; expected %d", got, expected)
		}
	}
}
