package main

import (
	"fmt"
	"strconv"
	"strings"
)

//
// https://www.hackerrank.com/challenges/time-conversion/problem
//

func solve(date string) string {
	var err error
	var hours uint64
	if hours, err = strconv.ParseUint(date[0:2], 10, 64); err == nil {
		isAM := strings.HasSuffix(date, "AM")

		if hours == 12 && isAM {
			hours = 00
		} else {
			if hours != 12 && !isAM {
				hours += 12
			}
		}

		return fmt.Sprintf("%02d:%s", hours, date[3:len(date)-2])
	}

	return ""
}

func main() {
	var date string
	fmt.Scanf("%v", &date)

	res := solve(date)

	fmt.Printf("%s", res)
}
