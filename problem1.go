package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum int
	sum = 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum = sum + i
		} else if i%5 == 0 {
			sum = sum + i
		}
	}

	fmt.Println("Sum of all multiples of 3 and 5 below 100: " + strconv.Itoa(sum))

}
