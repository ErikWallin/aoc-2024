package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 55312
	got := Run(`125 17`, 25)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
