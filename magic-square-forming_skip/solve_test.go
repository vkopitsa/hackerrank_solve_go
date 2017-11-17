package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve([]row{row{N0: 4, N1: 9, N2: 2}, row{N0: 3, N1: 5, N2: 7}, row{N0: 8, N1: 1, N2: 5}})
	result := 1
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}

func TestSolve2(t *testing.T) {
	solve := solve([]row{row{N0: 4, N1: 8, N2: 2}, row{N0: 4, N1: 5, N2: 7}, row{N0: 6, N1: 1, N2: 6}})
	result := 4
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}
