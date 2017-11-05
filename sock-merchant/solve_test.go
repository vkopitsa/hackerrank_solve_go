package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve([]uint{10, 20, 20, 10, 10, 30, 50, 10, 20})
	if solve != 3 {
		t.Error("Expected 3 got ", solve)
	}
}
