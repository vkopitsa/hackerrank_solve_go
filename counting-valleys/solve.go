//
// https://www.hackerrank.com/challenges/counting-valleys/problem
//
package main

import "fmt"

func main() {

	var count uint
	var steps string
	fmt.Scanf("%v\n%v", &count, &steps)

	solve := solve(steps)

	fmt.Printf("%d", solve)
}

func solve(steps string) uint {
	var count uint = 0
	var level int = 0

	for _, step := range steps {

		if step == 85 && level == -1 {
			count += 1
		}

		if step == 85 {
			level += 1
		} else {
			level += -1
		}
	}

	return count
}
