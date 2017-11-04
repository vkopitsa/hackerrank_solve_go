package main

import "testing"

func TestSolve(t *testing.T) {
	countMax, countMin := solve([]uint64{10, 5, 20, 20, 4, 5, 2, 25, 1})
	if countMax != 2 || countMin != 4 {
		t.Error("Expected 2 4 got ", countMax, countMin)
	}
}

func TestSolve2(t *testing.T) {
	countMax, countMin := solve([]uint64{3, 4, 21, 36, 10, 28, 35, 5, 24, 42})
	if countMax != 4 || countMin != 0 {
		t.Error("Expected 4 0 got ", countMax, countMin)
	}
}
