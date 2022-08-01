package main

func main() {
	// for problem 9, 	a   < b   <  c
	// that means 		a  != b   != c
	// but also 		a   + b   +  c = 1000
	// finally:			a^2 + b^2 =  c^2
	// there is only ONE solution that fits these 4 constraints

	// brute force method with time complexity O(n^3) would be to iterate each number and check the conditions

	// a faster method would be to subtract c from 1000. b from the difference. and the final difference be a.
	// that way a <= b <= c at least. but confirms that they all equal 1000.
	// if a = b or b = c, disregard the permutation and try another 3 numbers.

	// okay i brute forced it because i figured what the ranges would be to limit it from being worst case n^3.
	out:
	for c:=997; c>3;c--{
		for b:=499;b>2;b--{
			for a:=249;a>1;a--{
				if a + b + c == 1000 && a*a + b*b == c*c{
					println(a, b, c)
					println(a*b*c)
					break out
				}
			}
		}
	}
}
