package main

// takes int, returns triangle number
func nthTriangleNumber(x int)(ret int){
	var triangle	int
	triangle = (x*(x+1))/2
	return triangle
}

// takes int, returns list of its divisors
func listDivisors(x int)(ret []int){
	var list	[]int

	for i:=1; i<=x; i++{
		if x%i == 0{
			list = append(list, i)
		}
	}
	return list
}

func main() {

	for n:=0; len(listDivisors(nthTriangleNumber(n))) < 500; n++{
		println(nthTriangleNumber(n+1))
	}
}
