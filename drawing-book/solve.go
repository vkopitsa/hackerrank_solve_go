//
// https://www.hackerrank.com/challenges/drawing-book/problem
//
package main

import "fmt"

func main() {

	var count, page uint
	fmt.Scanf("%v\n%v", &count, &page)

	solve := solve(count, page)

	fmt.Printf("%d", solve)
}

func solve(count, page uint) uint {
	var pages uint = 0

	pages1 := ((count - 1) / 2) + 1
	if count%2 == 0 {
		pages1 = pages1 + 1
	}

	pages = ((page - 1) / 2) + 1
	if page%2 == 0 {
		pages = pages + 1
	}

	//fmt.Printf("%d - pages1, %d - pages\n", pages1, pages)

	if pages1/2 >= pages {
		pages = pages - 1
	} else {
		pages = pages1 - pages
	}

	return pages
}
