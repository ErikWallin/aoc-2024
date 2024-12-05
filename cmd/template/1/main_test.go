package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 0
	got := Run(``)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}
