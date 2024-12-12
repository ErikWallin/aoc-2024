package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 80
	got := Run(`AAAA
BBCD
BBCC
EEEC`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}

func TestRun2(t *testing.T) {
	expected := 436
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
	expected := 236
	got := Run(`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}

func TestRun4(t *testing.T) {
	expected := 368
	got := Run(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}

func TestRun5(t *testing.T) {
	expected := 1206
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
