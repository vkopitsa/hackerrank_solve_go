package main

import "testing"

func TestSolve(t *testing.T) {
	v := solve([]int{3, 2, 1, 3})
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
}
