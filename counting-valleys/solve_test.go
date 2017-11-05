package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve("UDDDUDUU")
	if solve != 1 {
		t.Error("Expected 1 got ", solve)
	}
}
