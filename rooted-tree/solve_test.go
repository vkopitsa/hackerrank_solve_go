package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	solve := solve([]row{row{A: 1, B: 2, K: 100}, row{A: 2, B: 5, K: 100}, row{A: 3, B: 4, K: 100}}, 5)
	var result int64 = 200
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}

func TestSolve2(t *testing.T) {
	file, err := os.Open("input10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rowCount, listSize int64
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &listSize, &rowCount)

	rows := make([]row, rowCount)
	var i int64 = 0
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%v %v %v", &rows[i].A, &rows[i].B, &rows[i].K)
		i++
	}

	solve := solve(rows, listSize)
	var result int64 = 2510535321
	if solve != result {
		t.Errorf("Expected %d got %d", result, solve)
	}
}
