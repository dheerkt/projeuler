package main

import "fmt"

func main() {

	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143 ?

	// num to prime factorize
	num := 600851475143

	// ub := num/2
	// r := []int{}

	/* Initial brute force solution, computationally inefficient */
	// for i := 2; i < ub; i++ {
	// 	if num % i == 0 {
	// 		num /= i
	// 		r = append(r, i)
	// 		fmt.Println(r)
	// 	}
	// }

	// fmt.Println(r[len(r)-1])

	i := 2
	for((i * i) < num) {
		for((num % i) == 0) {
			num /= i
		}
		i++
	}

	fmt.Println(num)
}