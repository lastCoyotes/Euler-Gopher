package main

import (
	"fmt"
	"strconv"
)

// recursive fibonacci function
func fibonacci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	//variables
	var i	int
	var sum int

	// init
	sum = 0
	i = 0

	for fibonacci(i) < 4000000 {
		if fibonacci(i)%2 == 0 {
			sum = fibonacci(i) + sum
		}
		i++
	}

	fmt.Println("Sum of all even Fibonacci numbers below 4 million: " + strconv.Itoa(sum))
}
