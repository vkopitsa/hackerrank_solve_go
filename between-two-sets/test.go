//
// https://www.hackerrank.com/challenges/between-two-sets/problem
//
// Sample Input
// 2 3
// 2 4
// 16 32 96
//
// Sample Output
// 3

package main

import "fmt"

func main() {

	var c1, c2 int
	fmt.Scanf("%v %v", &c1, &c2)

	firstArray := make([]int, c1)
	for i := range firstArray {
		_, err := fmt.Scan(&firstArray[i])
		if err != nil {
			firstArray = firstArray[:i]
			break
		}
	}

	secondArray := make([]int, c2)
	for i := range secondArray {
		_, err := fmt.Scan(&secondArray[i])
		if err != nil {
			secondArray = secondArray[:i]
			break
		}
	}

	count := solve(firstArray, secondArray)

	fmt.Printf("%d", count)
}

func solve(firstArray []int, secondArray []int) int {
	max := 100
	var sum int

	for i := 1; i <= max; i++ {
		isMod1 := true
		for _, ferstItem := range firstArray {
			if i%ferstItem != 0 {
				isMod1 = false
			}
		}
		isMod2 := true
		for _, secondItem := range secondArray {
			if secondItem%i != 0 {
				isMod2 = false
			}
		}

		if isMod1 && isMod2 {
			sum += 1
		}
	}

	return sum
}
