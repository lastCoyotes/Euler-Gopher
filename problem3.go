package main

import (
	"fmt"
	"strconv"
)

func main() {
	// init integers
	number := 600851475143
	i := 2

	//
	for {
		if number%i == 0{
			number = number / i
			fmt.Println("Current Factor " + strconv.Itoa(i))
		} else {
			i++
		}

		if i >= number{
			break
		}
	}

}
