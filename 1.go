package main

import "fmt"

func contains(i []int, num int) bool {
	for _, v := range i {
		if v == num {
			return true
		}
	}
	return false
}

func sum(i []int) int {
	r := 0
	for _, v := range i {
		r += v
	}
	return r
}

func main() {
	ub := 1000 // upper bound

	// find multiples of 3 and 5
	intArr := []int{3, 5}

	fmt.Printf("1. List all numbers below %d\n", ub)
	fmt.Printf("2. Find multiples of %v\n", intArr)
	fmt.Printf("3. Total all the multiples and print the sum\n\n")

	// list of multiples
	m := []int{}

	for i := 1; i < ub; i++ {
		for j := 0; j < len(intArr); j++ {
			if i%intArr[j] == 0 {
				if contains(m, i) != true {
					m = append(m, i)
				}
			}
		}
	}

	fmt.Printf("The sum is %d\n", sum(m))
}