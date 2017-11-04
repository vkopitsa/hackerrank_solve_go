package main

import "testing"

func TestSolve(t *testing.T) {
	date := solve(2017)
	if date != "13.09.2017" {
		t.Error("Expected 13.09.2017 got ", date)
	}
}

func TestSolve2(t *testing.T) {
	date := solve(2016)
	if date != "12.09.2016" {
		t.Error("Expected 12.09.2016 got ", date)
	}
}

func TestSolve3(t *testing.T) {
	date := solve(1918)
	if date != "26.09.1918" {
		t.Error("Expected 26.09.1918 got ", date)
	}
}
