//
// https://www.hackerrank.com/challenges/day-of-the-programmer/problem
//
package main

import "fmt"

func main() {

	var year uint
	fmt.Scanf("%v", &year)

	solve := solve(year)

	fmt.Printf("%s", solve)
}

func solve(year uint) string {
	day := "13"

	// is leap year of Julian and Gregorian calendar
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) || (year < 1918 && year%4 == 0) {
		day = "12"
	}

	if year == 1918 {
		day = "26"
	}

	return fmt.Sprintf("%s.09.%d", day, year)
}
