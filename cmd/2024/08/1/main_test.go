package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 14
	got := Run(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}