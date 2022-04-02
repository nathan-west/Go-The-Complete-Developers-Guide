package main

import (
	"fmt"
	"strconv"
)

func main() {
	sliceOfInts := []int{}
	i := 0
	for i < 11 {
		sliceOfInts = append(sliceOfInts, i)
		i++
	}

	for _, value := range sliceOfInts {
		oddOrEven := "odd"

		if value%2 == 0 {
			oddOrEven = "even"
		}

		fmt.Println(strconv.Itoa(value) + " is " + oddOrEven)
	}
}
