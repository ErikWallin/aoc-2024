package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 48
	got := Run(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
