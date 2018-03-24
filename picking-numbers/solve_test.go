package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve([]int{4, 6, 5, 3, 3, 1})
	result := 3
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}

func TestSolve2(t *testing.T) {
	solve := solve([]int{1, 2, 2, 3, 1, 2})
	result := 5
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}

func TestSolve3(t *testing.T) {
	solve := solve([]int{66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66})
	result := 100
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}

func TestSolve4(t *testing.T) {
	solve := solve([]int{4, 97, 5, 97, 97, 4, 97, 4, 97, 97, 97, 97, 4, 4, 5, 5, 97, 5, 97, 99, 4, 97, 5, 97, 97, 97, 5, 5, 97, 4, 5, 97, 97, 5, 97, 4, 97, 5, 4, 4, 97, 5, 5, 5, 4, 97, 97, 4, 97, 5, 4, 4, 97, 97, 97, 5, 5, 97, 4, 97, 97, 5, 4, 97, 97, 4, 97, 97, 97, 5, 4, 4, 97, 4, 4, 97, 5, 97, 97, 97, 97, 4, 97, 5, 97, 5, 4, 97, 4, 5, 97, 97, 5, 97, 5, 97, 5, 97, 97, 97})
	result := 50
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}

func TestSolve5(t *testing.T) {
	solve := solve([]int{4, 2, 3, 4, 4, 9, 98, 98, 3, 3, 3, 4, 2, 98, 1, 98, 98, 1, 1, 4, 98, 2, 98, 3, 9, 9, 3, 1, 4, 1, 98, 9, 9, 2, 9, 4, 2, 2, 9, 98, 4, 98, 1, 3, 4, 9, 1, 98, 98, 4, 2, 3, 98, 98, 1, 99, 9, 98, 98, 3, 98, 98, 4, 98, 2, 98, 4, 2, 1, 1, 9, 2, 4})
	result := 22
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}
