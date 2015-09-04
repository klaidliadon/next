#Next
The package calculates asynchronusly combinations and permutations of a series of value.

##Usage
Just use `next.NewCombination()` to get a channel of results.

	package main
	
	import (
		"fmt"
		"github.com/klaidliadon/next"
	)
	
	func main() {
		c := next.Combination{Base: []interface{}{
			"a", "b", "c",
		}}
		for v := range c.Results(2) {
			fmt.Println(v)
		}
	}

##Roadmap
The goal is to cover the 4 main cases of combinatronics:

- [x] Combinations
- [ ] Combinations with repetitions
- [ ] Permutations
- [ ] Permutations with repetitions