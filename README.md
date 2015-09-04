#Next [![Build Status](https://travis-ci.org/klaidliadon/next.svg?branch=master)](https://travis-ci.org/klaidliadon/next) [![GoDoc](http://godoc.org/github.com/klaidliadon/next?status.png)](http://godoc.org/github.com/klaidliadon/next) 
==============================

The package calculates asynchronously combinations and permutations of a collection of values.

##Usage
Just use create a new `Combination{...}` and get channel of results.

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