package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 142
	got := Run(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
