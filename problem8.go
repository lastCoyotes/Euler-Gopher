package main

import (
	"strconv"
	"strings"
)

func main() {
	// number
	var n string
	var n2 string
	n = "73167176531330624919225119674426574742355349194934969835203127745063262395783180169848018694788518438586156078911294949545950173795833195285320880551112540698747158523863050715693290963295227443043557668966489504452445231617318564030987111217223831136222989342338030813533627661428280644448664523874930358907296290491560440772390713810515859307960866701724271218839987979087922749219016997208880937766572733300105336788122023542180975125454059475224352584907711670556013604839586446706324415722155397536978179778461740649551492908625693219784686224828397224137565705605749026140797296865241453510047482166370484403199890008895243450658541227588666881164271714799244429282308634656748139191231628245861786645835912456652947654568284891288314260769004224219022671055626321111109370544217506941658960408071984038509624554443629812309878799272442849091888458015616609791913387549920052406368991256071760605886116467109405077541002256983155200055935"
	n2 = "72972571636269561882670428252483600823257530420752963450"
	var x			[]string
	var z 			[]string
	// split number into array of single digits
	// extra array string is needed
	x = strings.Split(n, "")
	z = strings.Split(n2, "")

	for u:=0;u<len(z);u++{
		x = append(x, z[u])
	}

	// need slice of the 13 adjacent digits
	var thirteen 	[]string
	var product		int
	var largest 	int
	var tempint		int

	product = 1
	tempint = 1
	largest = 0

	// take 2 on the problem solving
	for i:=0;i<len(x);i++{
		// in order of verisimilitude
		if len(thirteen) < 13 {
			thirteen = append(thirteen, x[i])
		}
		if len(thirteen) == 13 {
			// evaluate (multiply each digit)
			product = 1
			for j := 0; j < 13; j++ {
				tempint, _ = strconv.Atoi(thirteen[j])
				product *= tempint
			}
			// check if larger than last stored largest int
			if product > largest {
				largest = product
				println(largest)
			}
			
			// test printing
			//for a:=0;a<len(thirteen);a++ {
			//	print(thirteen[a])
			//}
			//println()

			// shift the first element out of the array now
			// reassign thirteen array to all values indexed one over
			_, thirteen = thirteen[0], thirteen[1:]

			// test printing
			//for a:=0;a<len(thirteen);a++ {
			//	print(thirteen[a])
			//}
			//println()
		}
	}

	//result
	print(largest)

}
