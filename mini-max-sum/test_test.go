package main

import "testing"

func TestSolve(t *testing.T) {
	v := solve(1, 2, 3, 4, 5)
	if v[0] != 10 || v[1] != 14 {
		t.Error("Expected [10, 14], got ", v)
	}
}

// func TestSolve1(t *testing.T) {
// 	//var v float64
// 	v := solve(0, 0, 0, 184467440, 184467441)
// 	if v[0] != 1000 || v[1] != 1400 {
// 		t.Error("Expected [1000, 1400], got ", v)
// 	}
// }
