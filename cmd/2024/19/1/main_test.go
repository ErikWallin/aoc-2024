package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 6
	got := Run(`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
