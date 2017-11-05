package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve([]uint64{3, 10, 2, 9}, 1, 12)
	if solve != "5" {
		t.Error("Expected 5 got ", solve)
	}
}

func TestSolve2(t *testing.T) {
	solve := solve([]uint64{3, 10, 2, 9}, 1, 7)
	if solve != "Bon Appetit" {
		t.Error("Expected Bon Appetit got ", solve)
	}
}
