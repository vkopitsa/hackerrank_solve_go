//
// https://www.hackerrank.com/challenges/magic-square-forming/problem
//
package main

import "fmt"

type row struct {
	N0 int
	N1 int
	N2 int
}

func main() {
	square := make([]row, 3)
	for i := range square {
		fmt.Scanf("%v %v %v", &square[i].N0, &square[i].N1, &square[i].N2)
	}

	solve := solve(square)

	fmt.Printf("%d", solve)
}

func solve(square []row) int {
	var cost int
	//var magicSum int = 15

	// for _, s := range square {
	// 	if s.N0+s.N1+s.N2 != magicSum {
	// 		if (s.N0 + s.N1 + s.N2) > magicSum {
	// 			cost += (s.N0 + s.N1 + s.N2) - magicSum
	// 		} else {
	// 			cost += magicSum - (s.N0 + s.N1 + s.N2)

	// 		}
	// 	}
	// }

	for _, s := range square {
		//cost += Abs((s.N0+s.N1+s.N2)+(square[i].N0+square[i].N1+square[i].N2)-30) / 2
		cost += (s.N0 + s.N1 + s.N2)
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
