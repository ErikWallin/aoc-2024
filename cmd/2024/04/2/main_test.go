package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 9
	got := Run(`.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
