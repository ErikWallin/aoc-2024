package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 161
	got := Run(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)
    if got != expected {
        t.Errorf("Run = %d; expected %d", got, expected)
    }
}
