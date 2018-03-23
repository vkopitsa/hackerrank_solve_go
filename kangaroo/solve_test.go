package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve(0, 3, 4, 2)
	if solve != "YES" {
		t.Error("Expected YES got ", solve)
	}
}
