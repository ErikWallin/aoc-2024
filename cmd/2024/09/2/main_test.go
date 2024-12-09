package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 2858
	got := Run(`2333133121414131402`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}
