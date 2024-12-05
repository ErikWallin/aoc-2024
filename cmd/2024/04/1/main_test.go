package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 18
	got := Run(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}
