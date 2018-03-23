//
// https://www.hackerrank.com/challenges/kangaroo/problem
//
package main

import "fmt"

func main() {

	var x1, v1, x2, v2 int
	fmt.Scanf("%v %v %v %v", &x1, &v1, &x2, &v2)

	solve := solve(x1, v1, x2, v2)

	fmt.Printf("%s", solve)
}

func solve(x1, v1, x2, v2 int) string {
	land1, land2 := x1, x2
	for i := 1; i <= 10000; i++ {

		land1 += v1
		land2 += v2

		if land1 == land2 {
			return "YES"
		}
	}

	return "NO"
}
