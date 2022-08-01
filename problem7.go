package main

import (
	"fmt"
	"strconv"
)

func IsPrime(x int)(ret bool){
	// corner cases
	if x <= 1 {
		return false
	}
	if x == 2 || x == 3 {
		return true
	}

	// below 5 cases for 2 and 3
	if x % 2 == 0 || x % 3 == 0 {
		return false
	}

	for i := 5; i * i <= x; i += 6 {
		if x % i == 0 || x % ( i + 2 ) == 0 {
			return false
		}
	}
	return true
}

func nthPrime(x int)(ret int){
	a := 2
	y := x
	for y > 0{
		if IsPrime(a){
			y--
		}
		a++
	}
	a--

	return a
}

func main() {
	
	fmt.Println("The 10,001th prime number is: " + strconv.Itoa(nthPrime(10001)))

}
