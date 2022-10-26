// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "fmt"

// generate array with numbers ranging from start to end
func makeRange(s int, e int) []int {
	r := make([]int, e-s+1)	// make empty array of size e-s+1
	for i := range r {
		r[i] = s+i
	}
	return r
}

func main() {
	c := makeRange(1, 20)
	i := 1
	notFound := true

	for(notFound) {
		i++
		for _, v := range c {
			if(i % v == 0) {
				if(v == c[len(c)-1]) {
					notFound = false
					break
				}
				continue
			} else {
				break
			}
		}
	}

	fmt.Println(i)
}