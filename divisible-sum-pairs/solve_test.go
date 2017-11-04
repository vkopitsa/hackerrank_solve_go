package main

import "testing"

func TestSolve(t *testing.T) {
	count := solve([]uint{1, 3, 2, 6, 1, 2}, 3)
	if count != 5 {
		t.Error("Expected 5 got ", count)
	}
}

