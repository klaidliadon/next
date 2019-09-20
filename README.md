# Next
[![GoDoc](https://godoc.org/klaidliadon.dev/next?status.svg)](https://godoc.org/klaidliadon.dev/next)
[![Build Status](https://travis-ci.org/klaidliadon/next.svg?branch=master)](https://travis-ci.org/klaidliadon/next) 
[![codecov.io](http://codecov.io/github/klaidliadon/next/coverage.svg?branch=master)](http://codecov.io/github/klaidliadon/next?branch=master)

The package calculates asynchronously combinations and permutations of a collection of values.

To add the package recommended:

	go get klaidliadon.dev/next

## Functionalities

- Combinations
- Combinations with repetitions
- Permutations
- Permutations with repetitions

## Usage

Just use create a new channel using `Combination()` or `Permutation()` and receive the results.

	package main
	
	import (
		"fmt"
		"klaidliadon.dev/next"
	)
	
	func main() {
		for v := range next.Combination([]interface{}{1,2,3,4}, 2, true) {
			fmt.Println(v)
		}
	}

Produces

	[1 1]
	[1 2]
	[1 3]
	[1 4]
	[2 2]
	[2 3]
	[2 4]
	[3 3]
	[3 4]
	[4 4]

## Roadmap

The 4 main cases of combinatronics are covered:
 
- [x] Combinations
- [x] Combinations with repetitions
- [x] Permutations
- [x] Permutations with repetitions

The future updates will improve performance and memory usage.
