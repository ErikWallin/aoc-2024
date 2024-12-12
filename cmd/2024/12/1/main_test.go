package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 140
	got := Run(`AAAA
BBCD
BBCC
EEEC`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}

func TestRun2(t *testing.T) {
	expected := 772
	got := Run(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}

func TestRun3(t *testing.T) {
	expected := 1930
	got := Run(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}