//
// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
//
package main

import "fmt"

func main() {

	var count, div uint
	fmt.Scanf("%v %v", &count, &div)

	array := make([]uint, count)
	for i := range array {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			array = array[:i]
			break
		}
	}

	solve := solve(array, div)

	fmt.Printf("%d", solve)
}

func solve(array []uint, div uint) int {
	var pairs int
	for i1, n1 := range array {
		for i2, n2 := range array {
			if i1 < i2 && (n1+n2)%div == 0 {
				pairs += 1
			}
		}
	}

	return pairs
}
