//
// https://www.hackerrank.com/challenges/migratory-birds/problem
//
package main

import "fmt"

func main() {

	var count uint64
	fmt.Scanf("%v", &count)

	array := make([]uint, count)
	for i := range array {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			array = array[:i]
			break
		}
	}

	solve := solve(array)

	fmt.Printf("%d", solve)
}

func solve(flock []uint) uint {
	birdTypes := map[uint]uint{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}

	for _, bird := range flock {
		_, ok := birdTypes[bird]
		if ok {
			birdTypes[bird] += 1
		}
	}

	var previousCount, birdType uint
	for _, t := range []uint{1, 2, 3, 4, 5} {
		countBirdType, ok := birdTypes[t]
		if ok && countBirdType > 0 && countBirdType > previousCount {
			previousCount = countBirdType
			birdType = t
		}
	}

	return birdType
}
