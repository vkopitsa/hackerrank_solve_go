package main

import "testing"

func TestSolve(t *testing.T) {
	mat := [][]int{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}}
	solve := solve(mat)
	result := 1
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}

func TestSolve2(t *testing.T) {
	mat := [][]int{{4, 8, 2}, {4, 5, 7}, {6, 1, 6}}
	solve := solve(mat)
	result := 4
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}

// Test Case #10
// 4 1 5 - 10 | 5
// 7 6 4 - 17 | 2
// 7 2 2 - 9  | 5
// - - - -    |12
// 18 9 11    |
// 3  6 4     |13
// 21

func TestSolve3(t *testing.T) {
	mat := [][]int{{4, 1, 5}, {7, 6, 4}, {7, 2, 2}}
	solve := solve(mat)
	result := 21
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}
