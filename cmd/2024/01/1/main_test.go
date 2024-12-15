package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 11
	got := Run(`3   4
4   3
2   5
1   3
3   9
3   3`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
