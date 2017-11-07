package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve(10, []uint{3, 1}, []uint{5, 2, 8})
	if solve != 9 {
		t.Error("Expected 9 got ", solve)
	}
}

func TestSolve2(t *testing.T) {
	solve := solve(5, []uint{4}, []uint{5})
	if solve != -1 {
		t.Error("Expected -1 got ", solve)
	}
}
