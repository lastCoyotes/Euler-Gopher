package main

import (
	"fmt"
	"strconv"
	"strings"
)

// checks if string of numbers is a palindrome
func checkpalindrome(x int) (ret bool) {
	// init variables
	// split string of number
	a := strconv.Itoa(x)
	b := strings.Split(a, "")
	c := ""

	// Reverse array
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	// Append to string c
	for k := 0; k < len(b); k++ {
		c = c + b[k]
	}

	// check
	if a == c{
		return true
	} else {
		return false
	}

}

func main() {
	var x, y, largest int
	largest = 0

	// time n^2 complexity
	for x = 999; x > 0; x--{
		for y = 999; y > 0; y--{
			// fmt.Println("Integers: " + strconv.Itoa(x) + ", " + strconv.Itoa(y))
			if checkpalindrome(x*y){
				if x*y > largest{
					largest = x*y
					fmt.Println("Integers: " + strconv.Itoa(x) + " and " + strconv.Itoa(y))
				} else {
					break
				}
			}
		}
	}

	fmt.Println("Largest palindrome from the product of two 3-digit numbers is " + strconv.Itoa(largest))

}
