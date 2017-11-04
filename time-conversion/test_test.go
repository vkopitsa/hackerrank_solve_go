package main

import "testing"

func TestSolve(t *testing.T) {
	v := solve("07:05:45PM")
	if v != "19:05:45" {
		t.Error("Expected 19:05:45, got ", v)
	}
}

func TestSolve2(t *testing.T) {
	v := solve("12:00:00AM")
	if v != "00:00:00" {
		t.Error("Expected 00:00:00, got ", v)
	}
}

func TestSolve3(t *testing.T) {
	v := solve("12:00:00PM")
	if v != "12:00:00" {
		t.Error("Expected 12:00:00, got ", v)
	}
}
