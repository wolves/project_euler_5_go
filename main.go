package main

import "fmt"

func main() {
	number := 2
	var smallestDivisble int
	fmt.Println("Calculating")
	for {
		number += 2
		for i := 1; i <= 20; i++ {
			if number%i == 0 {
				smallestDivisble = number
			} else {
				smallestDivisble = 0
				break
			}
		}

		if smallestDivisble != 0 {
			fmt.Println("We have a winner!", smallestDivisble)
			break
		}
	}
}

/*
2520 is the smallest number that can be divided by each of the numbers from 1
to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?
*/
