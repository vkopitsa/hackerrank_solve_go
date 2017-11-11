//
// https://www.hackerrank.com/challenges/crush/problem
//
package main

import "fmt"

type row struct {
	A int64
	B int64
	K int64
}

func main() {

	var listSize, rowCount int64
	fmt.Scanf("%v %v", &listSize, &rowCount)

	list := make([]row, rowCount)
	for i := range list {
		fmt.Scanf("%v %v %v", &list[i].A, &list[i].B, &list[i].K)
	}

	solve := solve(list, listSize)

	fmt.Printf("%d", solve)
}

func solve(rows []row, listSize int64) int64 {
	list := make([]int64, listSize+1)
	for _, r := range rows {
		list[r.A] += r.K
		if r.B+1 <= listSize {
			list[r.B+1] -= r.K
		}
	}

	var max, sum int64
	for _, v := range list {
		sum += v
		if max < sum {
			max = sum
		}
	}

	return max
}
