//Page 2
//Euclids Algorithm
//Given two positive integers m and n, find the largest positive integer that evenly devides both m and n

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	
	//Check input
	if len(os.Args) != 3 {
		fmt.Println("Supply two positive intergers, I will compute the greatest common divisor");
		return
	}


	//Large Integer
	m, err := strconv.Atoi(os.Args[1])

	//Small Integer
	n, err := strconv.Atoi(os.Args[2])

	//Remainder
	r := 0


	//Exit on input errors
	if err != nil {
		fmt.Println("Error converting your input into ints")
		fmt.Println(err.Error())
		return
	}

	

	//E0 - Ensure M >= N 
	//Swap values if n is larger than m
	if n > m {
		temp := n
		n = m
		m = temp
	}


	for {

		
		//E1 - Find Remainder, Loop
		r = m % n
		println("M:", m, " N:", n, "R:", r)
		//E2 - Is R Zero?
		if r == 0 {break}
		
		//E3 - Reduce
		m = n
		n = r
	}

	fmt.Println(n, " is the greatest common devisor");

	


	//Find Remainder
	//Is it Zeroconst
	//Reduce

	
}