//
// https://www.hackerrank.com/challenges/sock-merchant/problem
//
package main

import "fmt"

func main() {

	var count uint
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

func solve(socks []uint) uint {
	var pairs uint = 0
	ignore := make(map[int]int)
	for i1, sock1 := range socks {
		for i2, sock2 := range socks {
			_, ok2 := ignore[i2]
			_, ok1 := ignore[i1]
			if ok2 || ok1 || i1 == i2 {
				continue
			}
			if sock1 == sock2 {
				pairs += 1
				ignore[i1] = i1
				ignore[i2] = i2
			}
		}
	}

	return pairs
}
