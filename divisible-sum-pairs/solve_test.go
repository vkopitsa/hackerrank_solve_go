package main

import "testing"

func TestSolve(t *testing.T) {
	count := solve([]uint64{1, 2, 1, 3, 2}, 3, 2)
	if count != 2 {
		t.Error("Expected 2 got ", count)
	}
}

func TestSolve2(t *testing.T) {
	count := solve([]uint64{1, 1, 1, 1, 1, 1}, 3, 2)
	if count != 0 {
		t.Error("Expected 0 got ", count)
	}
}

func TestSolve3(t *testing.T) {
	count := solve([]uint64{4}, 4, 1)
	if count != 1 {
		t.Error("Expected 1 got ", count)
	}
}

func TestSolve4(t *testing.T) {
	count := solve([]uint64{0, 0, 0, 0, 0, 0, 0, 0}, 3, 2)
	if count != 0 {
		t.Error("Expected 0 got ", count)
	}
}

func TestSolve5(t *testing.T) {
	count := solve([]uint64{2, 2, 2, 2, 2, 2, 2, 2}, 4, 2)
	if count != 7 {
		t.Error("Expected 7 got ", count)
	}
}
