package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 4
	got := Run(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
