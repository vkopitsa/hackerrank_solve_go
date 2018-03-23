//
// https://www.hackerrank.com/challenges/magic-square-forming/problem
//
package main

import "fmt"

func main() {
	square := make([][]int, 3)
	for i := range square {
		square[i] = make([]int, 3)
		fmt.Scanf("%v %v %v", &square[i][0], &square[i][1], &square[i][2])
	}

	solve := solve(square)

	fmt.Printf("%d", solve)
}

func solve(square [][]int) int {
	cost := 81

	mat := [][][]int{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}

	for k := 0; k < 8; k++ {
		tmpCost := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				tmpCost += Abs(square[i][j] - mat[k][i][j])
			}
		}
		if tmpCost < cost {
			cost = tmpCost
		}
	}

	return cost
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}
