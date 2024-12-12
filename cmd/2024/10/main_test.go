package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 81
	got := Run(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}
