package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 16
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

func TestRun_bggr(t *testing.T) {
	expected := 1
	got := Run(`r, wr, b, g, bwu, rb, gb, br

bggr`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}

func TestRun_gbbr(t *testing.T) {
	expected := 4
	got := Run(`r, wr, b, g, bwu, rb, gb, br

gbbr`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}

func TestRun_rrbgbr(t *testing.T) {
	expected := 6
	got := Run(`r, wr, b, g, bwu, rb, gb, br

rrbgbr`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
