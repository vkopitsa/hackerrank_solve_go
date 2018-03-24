//
// https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
//
package main

import "fmt"
import "sort"

func main() {
	var count int
	fmt.Scanf("%v", &count)

	queries := make([]int, count)
	for i := range queries {
		fmt.Scanf("%v", &queries[i])
	}

	var count1 int
	fmt.Scanf("%v", &count1)

	queries1 := make([]int, count1)
	for i := range queries1 {
		fmt.Scanf("%v", &queries1[i])
	}

	solve := solve(queries, queries1)

	fmt.Printf("%d", solve)
}

func solve(queries []int, queries1 []int) []int {
	arr := make([]int, 200)
	for i := 0; i < len(queries); i++ {
		arr[queries[i]] = queries[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	integers := []int{}

	for _, d := range queries1 {

		for i := 0; i < len(arr); i++ {
			if d >= arr[i] {
				integers = append(integers, i+1)
				break
			}

		}
	}

	return integers
}
