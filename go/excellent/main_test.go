package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even" {
		t.Errorf("Expected: Even, Actual: %s", result)
	}
}