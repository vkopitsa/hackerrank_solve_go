//
// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem
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

	countMax, countMin := solve(array)

	fmt.Printf("%d %d", countMax, countMin)
}

func solve(scores []uint64) (uint64, uint64) {
	var countMax, countMin, previousMax, previousMin uint64

	previousMax = scores[0]
	previousMin = scores[0]

	for _, score := range scores[1:] {

		if score > previousMax {
			countMax += 1
			previousMax = score
		}

		if score < previousMin {
			countMin += 1
			previousMin = score
		}
	}

	return countMax, countMin
}
