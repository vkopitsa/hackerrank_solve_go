package main

import "testing"

func TestSolve(t *testing.T) {
	count := solve([]uint{1, 4, 4, 4, 5, 3})
	if count != 4 {
		t.Error("Expected 4 got ", count)
	}
}

func TestSolve2(t *testing.T) {
	count := solve([]uint{3, 3, 3, 5, 5, 5})
	if count != 3 {
		t.Error("Expected 3 got ", count)
	}
}
