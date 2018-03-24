//
// https://www.hackerrank.com/challenges/picking-numbers/problem
//
package main

import "fmt"

func main() {
	var count int
	fmt.Scanf("%v", &count)

	queries := make([]int, count)
	for i := range queries {
		fmt.Scanf("%v", &queries[i])
	}

	solve := solve(queries)

	fmt.Printf("%d", solve)
}

func solve(queries []int) int {
	arr := make([]int, 100)
	for i := 0; i < len(queries); i++ {
		arr[queries[i]]++
	}

	max := 0
	for i := 1; i < len(arr)-1; i++ {
		sum := arr[i] + arr[i+1]
		if sum > max {
			max = sum
		}
	}

	return max
}
