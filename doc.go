/*
	Package next calculates asynchronously combinations and permutations of a collection of values.

	Here is a sample usage of next:

		package main

		import (
			"fmt"
			"klaidliadon.dev/next"
		)

		func main() {
			for v := range next.Combination([]int{1,2,3,4}, 2, true) {
				fmt.Println(v)
			}
		}
*/
package next
