//
// https://www.hackerrank.com/challenges/the-birthday-bar/problem
//

package main

import "fmt"

func main() {

	var c int
	fmt.Scanf("%v", &c)

	array := make([]uint64, c)
	for i := range array {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			array = array[:i]
			break
		}
	}

	var d, m uint64
	fmt.Scanf("%v %v", &d, &m)

	count := solve(array, d, m)

	fmt.Printf("%d", count)
}

func solve(squares []uint64, d, m uint64) int {
	var ways int
	var sum uint64
	var squaresLen = uint64(len(squares))

	if m <= squaresLen {
		for i := 0; uint64(i) < m; i++ {
			sum += squares[i]
		}

		if sum == d {
			ways += 1
		}
	}

	for i := 0; uint64(i) < squaresLen-m; i++ {
		sum = sum - squares[i] + squares[uint64(i)+m]
		if sum == d {
			ways++
		}
	}

	return ways
}
