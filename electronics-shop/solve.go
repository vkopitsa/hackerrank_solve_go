//
// https://www.hackerrank.com/challenges/electronics-shop/problem
//
package main

import "fmt"

func main() {

	var amount, keyboardCount, driveCount uint
	fmt.Scanf("%v %v %v", &amount, &keyboardCount, &driveCount)

	keyboardsPrice := make([]uint, keyboardCount)
	for i := range keyboardsPrice {
		_, err := fmt.Scan(&keyboardsPrice[i])
		if err != nil {
			keyboardsPrice = keyboardsPrice[:i]
			break
		}
	}

	drivesPrice := make([]uint, driveCount)
	for i := range drivesPrice {
		_, err := fmt.Scan(&drivesPrice[i])
		if err != nil {
			drivesPrice = drivesPrice[:i]
			break
		}
	}

	solve := solve(amount, keyboardsPrice, drivesPrice)

	fmt.Printf("%d", solve)
}

func solve(amount uint, keyboardsPrice []uint, drivesPrice []uint) int {
	var sum int = -1
	for _, keyboardPrice := range keyboardsPrice {
		for _, drivePrice := range drivesPrice {
			if (keyboardPrice + drivePrice) <= amount {
				if int(keyboardPrice+drivePrice) > sum {
					sum = int(keyboardPrice + drivePrice)
				}
			}
		}
	}

	return sum
}
