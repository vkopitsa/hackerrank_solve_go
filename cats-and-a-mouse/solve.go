//
// https://www.hackerrank.com/challenges/cats-and-a-mouse/problem
//
package main

import "fmt"

type query struct {
	CatA  int
	CatB  int
	Mouse int
}

func main() {
	var count uint
	fmt.Scanf("%v", &count)

	queries := make([]query, count)
	for i := range queries {
		fmt.Scanf("%v %v %v", &queries[i].CatA, &queries[i].CatB, &queries[i].Mouse)
	}

	solve := solve(queries)

	fmt.Printf("%s", solve)
}

func solve(queries []query) string {
	var result = ""
	for _, query := range queries {
		sa := Abs(query.Mouse - query.CatA)
		sb := Abs(query.Mouse - query.CatB)

		if result != "" {
			result += "\n"
		}

		if sa == sb {
			result += "Mouse C"
		} else if sa > sb {
			result += "Cat B"
		} else {
			result += "Cat A"
		}
	}

	return result
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
