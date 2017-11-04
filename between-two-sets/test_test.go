package main

import "testing"

func TestSolve(t *testing.T) {
	v := solve([]int{2, 4}, []int{16, 32, 96})
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}
