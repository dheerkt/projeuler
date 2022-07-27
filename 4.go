package main

import "fmt"

// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

func checkPal (n int) bool {
	res := 0
	for n > 999 {
		r := n % 10
		res = (res*10)+r
		n /= 10
	}

	if(n == res) {
		return true
	} else {
		return false
	}
}

func main() {
	// Flaws: the 900 lower bound for i and j is arbitrary. Got lucky there.
	// Another way to do it is to compute all the palindromes from i=900-999 and j = 900-999 and then pick the largest
	// That guarantees the largest but you have to do more computation.
	res := 0
	for i := 999; i > 900; i-- {
		for j := 999; j > 900; j-- {
			if(checkPal(i*j)) {
				res = i*j
				fmt.Println("Found! ", i, "*", j, "=", res)
				break
			}
		}
		if(res != 0) {
			break
		}
		
	}

	fmt.Println("Done!")
}

