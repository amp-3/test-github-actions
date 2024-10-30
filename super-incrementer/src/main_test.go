package main

import (
	"strconv"
	"testing"
)

func TestIncrement(t *testing.T) {
	result := Increment(2)
	if result != 3 {
		t.Errorf("expected: Increment() value: %s", strconv.Itoa(result))
	}
}
