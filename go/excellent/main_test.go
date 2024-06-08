package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even Number" {
		t.Errorf("Expected: Even Number, Actual: %s", result)
	}
}