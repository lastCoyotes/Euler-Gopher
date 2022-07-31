package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i			int

	i = 1
	for {
		if i%20 == 0 && i%19 == 0 && i%18 == 0 && i%17 == 0 && i%16 == 0 && i%15 == 0 && i%14 == 0 && i%13 == 0 && i%12 == 0 && i%11== 0{
			break
		} else {
			i++
		}
	}

	fmt.Println("Smallest multiple of all numbers 1-20 is: " + strconv.Itoa(i))

}
