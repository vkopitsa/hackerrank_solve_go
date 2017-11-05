package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve(6, 2)
	if solve != 1 {
		t.Error("Expected 1 got ", solve)
	}
}

func TestSolve2(t *testing.T) {
	solve := solve(5, 4)
	if solve != 0 {
		t.Error("Expected 0 got ", solve)
	}
}
