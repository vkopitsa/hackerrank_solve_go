package main

import "testing"

func TestSolve(t *testing.T) {
	solve := solve([]query{query{CatA: 1, CatB: 2, Mouse: 3}, query{CatA: 1, CatB: 3, Mouse: 2}, query{CatA: 2, CatB: 1, Mouse: 3}})
	result := `Cat B
Mouse C
Cat A`
	if solve != result {
		t.Errorf("Expected \n%s \ngot \n%s", result, solve)
	}
}
