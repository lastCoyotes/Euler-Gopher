package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i				int
	var difference		int
	var sumofsquares	int
	var squaredsum		int

	for i = 0; i < 101; i++{
		sumofsquares += i*i
		squaredsum += i
	}

	squaredsum *= squaredsum
	difference = squaredsum - sumofsquares

	fmt.Println("The difference between the sum of the squares of the first one hundred natural numbers and the square of the sum is: " + strconv.Itoa(difference))

}
