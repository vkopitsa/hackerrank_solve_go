package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	solve := solve([]int{100, 100, 50, 40, 40, 20, 10}, []int{5, 25, 50, 120})
	result := []int{6, 4, 2, 1}
	if !reflect.DeepEqual(solve, result) {
		t.Errorf("Expected %d got %d", result, solve)
	}
}
