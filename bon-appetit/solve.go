//
// https://www.hackerrank.com/challenges/bon-appetit/problem
//
package main

import "fmt"

func main() {

	var count, item uint64
	fmt.Scanf("%v %v", &count, &item)

	array := make([]uint64, count)
	for i := range array {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			array = array[:i]
			break
		}
	}

	var cost uint64
	fmt.Scanf("%v", &cost)

	solve := solve(array, item, cost)

	fmt.Printf("%s", solve)
}

func solve(items []uint64, item, cost uint64) string {
	solve := "Bon Appetit"
	var sum uint64 = 0
	for i, v := range items {
		if uint64(i) == item {
			continue
		}

		sum += v
	}

	if sum/2 != cost {
		solve = fmt.Sprintf("%d", cost-(sum/2))
	}

	return solve
}
