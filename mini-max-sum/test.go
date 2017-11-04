package main

import "fmt"

//
// https://www.hackerrank.com/challenges/mini-max-sum/problem
//
//

func solve(n1 uint64, n2 uint64, n3 uint64, n4 uint64, n5 uint64) []uint64 {
	var sum uint64 = 0
	for _, i := range []uint64{n1, n2, n3, n4, n5} {
		sum += i
	}

	var max uint64 = 0
	var min uint64 = 0
	for _, value := range []uint64{n1, n2, n3, n4, n5} {
		if max < value {
			max = value
		}

		if min == 0 || min > value {
			min = value
		}
	}

	return []uint64{sum - max, sum - min}
}

func main() {
	var n1, n2, n3, n4, n5 uint64
	fmt.Scanf("%v %v %v %v %v", &n1, &n2, &n3, &n4, &n5)
	res := solve(n1, n2, n3, n4, n5)

	fmt.Printf("%d %d", res[0], res[1])
}
