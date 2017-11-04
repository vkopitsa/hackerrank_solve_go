package main

import "fmt"

//
// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
//

func solve(candles []int) int {
	var max int = 0
	for _, value := range candles {
		if max < value {
			max = value
		}
	}

	var sum int = 0
	for _, value := range candles {
		if max == value {
			sum += 1
		}
	}

	return sum
}

func main() {

	var n int
	fmt.Scanf("%v", &n)

	candles := make([]int, n)
	for i := range candles {
		_, err := fmt.Scan(&candles[i])
		if err != nil {
			candles = candles[:i]
			break
		}
	}

	res := solve(candles)

	fmt.Printf("%d", res)
}
